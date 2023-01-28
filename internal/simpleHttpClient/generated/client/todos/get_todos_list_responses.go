// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"test-go-pipeline/internal/simpleHttpClient/generated/models"
)

// GetTodosListReader is a Reader for the GetTodosList structure.
type GetTodosListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTodosListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTodosListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTodosListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTodosListOK creates a GetTodosListOK with default headers values
func NewGetTodosListOK() *GetTodosListOK {
	return &GetTodosListOK{}
}

/* GetTodosListOK describes a response with status code 200, with default header values.

OK
*/
type GetTodosListOK struct {
	Payload []*models.Item
}

func (o *GetTodosListOK) Error() string {
	return fmt.Sprintf("[GET /todos][%d] getTodosListOK  %+v", 200, o.Payload)
}
func (o *GetTodosListOK) GetPayload() []*models.Item {
	return o.Payload
}

func (o *GetTodosListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTodosListNotFound creates a GetTodosListNotFound with default headers values
func NewGetTodosListNotFound() *GetTodosListNotFound {
	return &GetTodosListNotFound{}
}

/* GetTodosListNotFound describes a response with status code 404, with default header values.

error
*/
type GetTodosListNotFound struct {
	Payload models.Error
}

func (o *GetTodosListNotFound) Error() string {
	return fmt.Sprintf("[GET /todos][%d] getTodosListNotFound  %+v", 404, o.Payload)
}
func (o *GetTodosListNotFound) GetPayload() models.Error {
	return o.Payload
}

func (o *GetTodosListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}