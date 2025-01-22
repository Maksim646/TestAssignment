// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/Maksim646/TestAssignment/internal/api/definition"
)

// GetSongVerseReader is a Reader for the GetSongVerse structure.
type GetSongVerseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSongVerseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSongVerseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSongVerseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSongVerseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /song/{id}/verse] GetSongVerse", response, response.Code())
	}
}

// NewGetSongVerseOK creates a GetSongVerseOK with default headers values
func NewGetSongVerseOK() *GetSongVerseOK {
	return &GetSongVerseOK{}
}

/*
GetSongVerseOK describes a response with status code 200, with default header values.

Get Song Verse Response
*/
type GetSongVerseOK struct {
	Payload *models.SongVersesResponse
}

// IsSuccess returns true when this get song verse o k response has a 2xx status code
func (o *GetSongVerseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get song verse o k response has a 3xx status code
func (o *GetSongVerseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get song verse o k response has a 4xx status code
func (o *GetSongVerseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get song verse o k response has a 5xx status code
func (o *GetSongVerseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get song verse o k response a status code equal to that given
func (o *GetSongVerseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get song verse o k response
func (o *GetSongVerseOK) Code() int {
	return 200
}

func (o *GetSongVerseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /song/{id}/verse][%d] getSongVerseOK %s", 200, payload)
}

func (o *GetSongVerseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /song/{id}/verse][%d] getSongVerseOK %s", 200, payload)
}

func (o *GetSongVerseOK) GetPayload() *models.SongVersesResponse {
	return o.Payload
}

func (o *GetSongVerseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SongVersesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSongVerseBadRequest creates a GetSongVerseBadRequest with default headers values
func NewGetSongVerseBadRequest() *GetSongVerseBadRequest {
	return &GetSongVerseBadRequest{}
}

/*
GetSongVerseBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSongVerseBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get song verse bad request response has a 2xx status code
func (o *GetSongVerseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get song verse bad request response has a 3xx status code
func (o *GetSongVerseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get song verse bad request response has a 4xx status code
func (o *GetSongVerseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get song verse bad request response has a 5xx status code
func (o *GetSongVerseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get song verse bad request response a status code equal to that given
func (o *GetSongVerseBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get song verse bad request response
func (o *GetSongVerseBadRequest) Code() int {
	return 400
}

func (o *GetSongVerseBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /song/{id}/verse][%d] getSongVerseBadRequest %s", 400, payload)
}

func (o *GetSongVerseBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /song/{id}/verse][%d] getSongVerseBadRequest %s", 400, payload)
}

func (o *GetSongVerseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSongVerseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSongVerseInternalServerError creates a GetSongVerseInternalServerError with default headers values
func NewGetSongVerseInternalServerError() *GetSongVerseInternalServerError {
	return &GetSongVerseInternalServerError{}
}

/*
GetSongVerseInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSongVerseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get song verse internal server error response has a 2xx status code
func (o *GetSongVerseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get song verse internal server error response has a 3xx status code
func (o *GetSongVerseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get song verse internal server error response has a 4xx status code
func (o *GetSongVerseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get song verse internal server error response has a 5xx status code
func (o *GetSongVerseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get song verse internal server error response a status code equal to that given
func (o *GetSongVerseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get song verse internal server error response
func (o *GetSongVerseInternalServerError) Code() int {
	return 500
}

func (o *GetSongVerseInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /song/{id}/verse][%d] getSongVerseInternalServerError %s", 500, payload)
}

func (o *GetSongVerseInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /song/{id}/verse][%d] getSongVerseInternalServerError %s", 500, payload)
}

func (o *GetSongVerseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSongVerseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
