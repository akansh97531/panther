package api

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/panther-labs/panther/api/lambda/database/models"
)

// FIXME: consider adding as stand-alone lambda to de-couple this from Athena api

// https://docs.aws.amazon.com/appsync/latest/devguide/tutorial-local-resolvers.html
// https://www.youtube.com/watch?v=QabeZ4sag44
// https://www.youtube.com/watch?v=UK-c1NkdLA0
// https://www.youtube.com/watch?v=nAYikoFCN4k

const (
	// what we send to appsync
	mutationTemplate = `
mutation {
     queryDone(input: {
      userData: "%s",
      queryId: "%s",
      workflowId: "%s"
     }) {
       userData
       queryId
       workflowId
     }
}
`
)

type GraphQlQuery struct {
	Query string `json:"query"` // when we marshal, this will escape the mutation JSON as required by graphQL
}

type GraphQlResponse struct {
	Data   interface{}   `json:"data"`
	Errors []interface{} `json:"errors"`
}

func (API) NotifyAppSync(input *models.NotifyAppSyncInput) (*models.NotifyAppSyncOutput, error) {
	output := &models.NotifyAppSyncOutput{}

	var err error
	defer func() {
		if err != nil {
			err = apiError(err) // lambda failed
		}
	}()

	// make sigv4 https request to appsync endpoint notifying query is complete, sending  queryId and workflowId
	appSyncEndpoint := os.Getenv("GRAPHQL_ENDPOINT")
	httpClient := http.Client{}
	signer := v4.NewSigner(awsSession.Config.Credentials)

	mutation := &GraphQlQuery{
		Query: fmt.Sprintf(mutationTemplate, input.UserDataToken, input.QueryID, input.WorkflowIdentifier),
	}
	jsonMessage, err := jsoniter.Marshal(mutation)
	if err != nil {
		err = errors.Wrapf(err, "json marshal failed for: %#v", input)
		return output, err
	}

	body := bytes.NewReader(jsonMessage) // JSON envelope for graphQL

	req, err := http.NewRequest("POST", appSyncEndpoint, body)
	if err != nil {
		err = errors.Wrapf(err, "new htttp request failed for: %#v", input)
		return output, err
	}
	req.Header.Add("Content-Type", "application/json")

	if awsSession.Config.Region == nil {
		err = errors.Wrapf(err, "failed to get aws region %#v", input)
		return output, err
	}

	_, err = signer.Sign(req, body, "appsync", *awsSession.Config.Region, time.Now().UTC())
	if err != nil {
		err = errors.Wrapf(err, "failed to v4 sign %#v", input)
		return output, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		err = errors.Wrapf(err, "failed to POST %#v", req)
		return output, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		err = errors.Errorf("failed to POST (%d): %s", resp.StatusCode, string(respBody))
		return output, err
	}

	graphQlResp := &GraphQlResponse{}
	err = jsoniter.Unmarshal(respBody, graphQlResp)
	if err != nil {
		err = errors.Wrapf(err, "json marshal failed for: %#v", string(respBody))
		return output, err
	}

	if len(graphQlResp.Errors) > 0 {
		err = errors.Errorf("graphQL error for %#v: %#v", input, graphQlResp)
		return output, err
	}

	return output, nil
}