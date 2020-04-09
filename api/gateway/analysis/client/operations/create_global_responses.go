// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// CreateGlobalReader is a Reader for the CreateGlobal structure.
type CreateGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGlobalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGlobalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateGlobalConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGlobalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGlobalCreated creates a CreateGlobalCreated with default headers values
func NewCreateGlobalCreated() *CreateGlobalCreated {
	return &CreateGlobalCreated{}
}

/*CreateGlobalCreated handles this case with default header values.

Global created successfully
*/
type CreateGlobalCreated struct {
	Payload *models.Global
}

func (o *CreateGlobalCreated) Error() string {
	return fmt.Sprintf("[POST /global][%d] createGlobalCreated  %+v", 201, o.Payload)
}

func (o *CreateGlobalCreated) GetPayload() *models.Global {
	return o.Payload
}

func (o *CreateGlobalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGlobalBadRequest creates a CreateGlobalBadRequest with default headers values
func NewCreateGlobalBadRequest() *CreateGlobalBadRequest {
	return &CreateGlobalBadRequest{}
}

/*CreateGlobalBadRequest handles this case with default header values.

Bad request
*/
type CreateGlobalBadRequest struct {
	Payload *models.Error
}

func (o *CreateGlobalBadRequest) Error() string {
	return fmt.Sprintf("[POST /global][%d] createGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *CreateGlobalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGlobalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGlobalConflict creates a CreateGlobalConflict with default headers values
func NewCreateGlobalConflict() *CreateGlobalConflict {
	return &CreateGlobalConflict{}
}

/*CreateGlobalConflict handles this case with default header values.

Global with the given ID already exists
*/
type CreateGlobalConflict struct {
}

func (o *CreateGlobalConflict) Error() string {
	return fmt.Sprintf("[POST /global][%d] createGlobalConflict ", 409)
}

func (o *CreateGlobalConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGlobalInternalServerError creates a CreateGlobalInternalServerError with default headers values
func NewCreateGlobalInternalServerError() *CreateGlobalInternalServerError {
	return &CreateGlobalInternalServerError{}
}

/*CreateGlobalInternalServerError handles this case with default header values.

Internal server error
*/
type CreateGlobalInternalServerError struct {
}

func (o *CreateGlobalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /global][%d] createGlobalInternalServerError ", 500)
}

func (o *CreateGlobalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
