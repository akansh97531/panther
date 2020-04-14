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
	"fmt"

	"github.com/aws/aws-sdk-go/service/lambda"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/panther-labs/panther/api/lambda/database/models"
)

// called by Step workflow to execute callback lambda after query has finished
func (API) InvokeNotifyLambda(input *models.InvokeNotifyLambdaInput) (*models.InvokeNotifyLambdaInput, error) {
	output := &models.InvokeNotifyLambdaInput{}

	var err error
	defer func() {
		if err != nil {
			err = apiError(err) // lambda failed
		}
	}()

	// these lambdas are expected to take userData, queryId and workflowId in as arguments
	var notifyInput models.NotifyInput
	notifyInput.QueryID = input.QueryID
	notifyInput.WorkflowIdentifier = input.WorkflowIdentifier
	notifyInput.UserDataToken = input.UserDataToken
	payload, err := jsoniter.MarshalToString(&notifyInput)
	if err != nil {
		err = errors.Wrapf(err, "failed to marshal %#v", input)
		return output, err
	}

	resp, err := lambdaClient.Invoke(&lambda.InvokeInput{
		FunctionName: &input.LambdaName,
		Payload:      []byte(fmt.Sprintf(`{ "%s": %s}`, input.MethodName, payload)), // genericapi
	})
	if err != nil {
		err = errors.Wrapf(err, "failed to invoke %#v", input)
		return output, err
	}

	if resp.FunctionError != nil {
		err = errors.Errorf("%s: failed to invoke %#v", *resp.FunctionError, input)
		return output, err
	}

	return output, nil
}