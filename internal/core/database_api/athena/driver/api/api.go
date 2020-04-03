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
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/glue"

	"github.com/panther-labs/panther/api/lambda/database/models"
	"github.com/panther-labs/panther/internal/core/database_api/athena/driver"
)

var (
	sess                = session.Must(session.NewSession())
	glueClient          = glue.New(sess)
	athenaClient        = athena.New(sess)
	athenaS3ResultsPath *string
)

func init() {
	if os.Getenv("ATHENA_BUCKET") != "" {
		results := "s3://" + os.Getenv("ATHENA_BUCKET") + "/athena_api/"
		athenaS3ResultsPath = &results
	}
}

// API provides receiver methods for each route handler.
type API struct{}

func (API) GetDatabases(input *models.GetDatabasesInput) (*models.GetDatabasesOutput, error) {
	return driver.GetDatabases(glueClient, input)
}

func (API) GetTables(input *models.GetTablesInput) (*models.GetTablesOutput, error) {
	return driver.GetTables(glueClient, input)
}

func (API) GetTablesDetail(input *models.GetTablesDetailInput) (*models.GetTablesDetailOutput, error) {
	return driver.GetTablesDetail(glueClient, input)
}

func (API) ExecuteQuery(input *models.ExecuteQueryInput) (*models.ExecuteQueryOutput, error) {
	return driver.ExecuteQuery(athenaClient, input, athenaS3ResultsPath)
}

func (API) ExecuteAsyncQuery(input *models.ExecuteAsyncQueryInput) (*models.ExecuteAsyncQueryOutput, error) {
	return driver.ExecuteAsyncQuery(athenaClient, input, athenaS3ResultsPath)
}

func (API) GetQueryStatus(input *models.GetQueryStatusInput) (*models.GetQueryStatusOutput, error) {
	return driver.GetQueryStatus(athenaClient, input)
}

func (API) GetQueryResults(input *models.GetQueryResultsInput) (*models.GetQueryResultsOutput, error) {
	return driver.GetQueryResults(athenaClient, input)
}
