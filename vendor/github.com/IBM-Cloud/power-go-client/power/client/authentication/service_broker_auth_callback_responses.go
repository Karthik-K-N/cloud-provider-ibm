// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerAuthCallbackReader is a Reader for the ServiceBrokerAuthCallback structure.
type ServiceBrokerAuthCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewServiceBrokerAuthCallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerAuthCallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerAuthCallbackOK creates a ServiceBrokerAuthCallbackOK with default headers values
func NewServiceBrokerAuthCallbackOK() *ServiceBrokerAuthCallbackOK {
	return &ServiceBrokerAuthCallbackOK{}
}

/* ServiceBrokerAuthCallbackOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthCallbackOK struct {
	Payload *models.AccessToken
}

func (o *ServiceBrokerAuthCallbackOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback][%d] serviceBrokerAuthCallbackOK  %+v", 200, o.Payload)
}
func (o *ServiceBrokerAuthCallbackOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *ServiceBrokerAuthCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthCallbackUnauthorized creates a ServiceBrokerAuthCallbackUnauthorized with default headers values
func NewServiceBrokerAuthCallbackUnauthorized() *ServiceBrokerAuthCallbackUnauthorized {
	return &ServiceBrokerAuthCallbackUnauthorized{}
}

/* ServiceBrokerAuthCallbackUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerAuthCallbackUnauthorized struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthCallbackUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback][%d] serviceBrokerAuthCallbackUnauthorized  %+v", 401, o.Payload)
}
func (o *ServiceBrokerAuthCallbackUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthCallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthCallbackInternalServerError creates a ServiceBrokerAuthCallbackInternalServerError with default headers values
func NewServiceBrokerAuthCallbackInternalServerError() *ServiceBrokerAuthCallbackInternalServerError {
	return &ServiceBrokerAuthCallbackInternalServerError{}
}

/* ServiceBrokerAuthCallbackInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthCallbackInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthCallbackInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/callback][%d] serviceBrokerAuthCallbackInternalServerError  %+v", 500, o.Payload)
}
func (o *ServiceBrokerAuthCallbackInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthCallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
