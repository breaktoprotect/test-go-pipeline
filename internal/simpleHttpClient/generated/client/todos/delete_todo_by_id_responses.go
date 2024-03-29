// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"test-go-pipeline/internal/simplehttpclient/generated/models"
)

// DeleteTodoByIDReader is a Reader for the DeleteTodoByID structure.
type DeleteTodoByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTodoByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTodoByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTodoByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTodoByIDOK creates a DeleteTodoByIDOK with default headers values
func NewDeleteTodoByIDOK() *DeleteTodoByIDOK {
	return &DeleteTodoByIDOK{}
}

/* DeleteTodoByIDOK describes a response with status code 200, with default header values.

deleteTodoById
*/
type DeleteTodoByIDOK struct {
	Payload *models.Item
}

func (o *DeleteTodoByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /todos/{id}][%d] deleteTodoByIdOK  %+v", 200, o.Payload)
}
func (o *DeleteTodoByIDOK) GetPayload() *models.Item {
	return o.Payload
}

func (o *DeleteTodoByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTodoByIDNotFound creates a DeleteTodoByIDNotFound with default headers values
func NewDeleteTodoByIDNotFound() *DeleteTodoByIDNotFound {
	return &DeleteTodoByIDNotFound{}
}

/* DeleteTodoByIDNotFound describes a response with status code 404, with default header values.

error
*/
type DeleteTodoByIDNotFound struct {
	Payload models.Error
}

func (o *DeleteTodoByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /todos/{id}][%d] deleteTodoByIdNotFound  %+v", 404, o.Payload)
}
func (o *DeleteTodoByIDNotFound) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteTodoByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
