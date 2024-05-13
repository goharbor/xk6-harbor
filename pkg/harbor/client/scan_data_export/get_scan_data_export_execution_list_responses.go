// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// GetScanDataExportExecutionListReader is a Reader for the GetScanDataExportExecutionList structure.
type GetScanDataExportExecutionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanDataExportExecutionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanDataExportExecutionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScanDataExportExecutionListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScanDataExportExecutionListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScanDataExportExecutionListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScanDataExportExecutionListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /export/cve/executions] getScanDataExportExecutionList", response, response.Code())
	}
}

// NewGetScanDataExportExecutionListOK creates a GetScanDataExportExecutionListOK with default headers values
func NewGetScanDataExportExecutionListOK() *GetScanDataExportExecutionListOK {
	return &GetScanDataExportExecutionListOK{}
}

/*
GetScanDataExportExecutionListOK describes a response with status code 200, with default header values.

Success
*/
type GetScanDataExportExecutionListOK struct {
	Payload *models.ScanDataExportExecutionList
}

// IsSuccess returns true when this get scan data export execution list o k response has a 2xx status code
func (o *GetScanDataExportExecutionListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scan data export execution list o k response has a 3xx status code
func (o *GetScanDataExportExecutionListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution list o k response has a 4xx status code
func (o *GetScanDataExportExecutionListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scan data export execution list o k response has a 5xx status code
func (o *GetScanDataExportExecutionListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution list o k response a status code equal to that given
func (o *GetScanDataExportExecutionListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get scan data export execution list o k response
func (o *GetScanDataExportExecutionListOK) Code() int {
	return 200
}

func (o *GetScanDataExportExecutionListOK) Error() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListOK  %+v", 200, o.Payload)
}

func (o *GetScanDataExportExecutionListOK) String() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListOK  %+v", 200, o.Payload)
}

func (o *GetScanDataExportExecutionListOK) GetPayload() *models.ScanDataExportExecutionList {
	return o.Payload
}

func (o *GetScanDataExportExecutionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScanDataExportExecutionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionListUnauthorized creates a GetScanDataExportExecutionListUnauthorized with default headers values
func NewGetScanDataExportExecutionListUnauthorized() *GetScanDataExportExecutionListUnauthorized {
	return &GetScanDataExportExecutionListUnauthorized{}
}

/*
GetScanDataExportExecutionListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetScanDataExportExecutionListUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution list unauthorized response has a 2xx status code
func (o *GetScanDataExportExecutionListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution list unauthorized response has a 3xx status code
func (o *GetScanDataExportExecutionListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution list unauthorized response has a 4xx status code
func (o *GetScanDataExportExecutionListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan data export execution list unauthorized response has a 5xx status code
func (o *GetScanDataExportExecutionListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution list unauthorized response a status code equal to that given
func (o *GetScanDataExportExecutionListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get scan data export execution list unauthorized response
func (o *GetScanDataExportExecutionListUnauthorized) Code() int {
	return 401
}

func (o *GetScanDataExportExecutionListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanDataExportExecutionListUnauthorized) String() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanDataExportExecutionListUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionListForbidden creates a GetScanDataExportExecutionListForbidden with default headers values
func NewGetScanDataExportExecutionListForbidden() *GetScanDataExportExecutionListForbidden {
	return &GetScanDataExportExecutionListForbidden{}
}

/*
GetScanDataExportExecutionListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScanDataExportExecutionListForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution list forbidden response has a 2xx status code
func (o *GetScanDataExportExecutionListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution list forbidden response has a 3xx status code
func (o *GetScanDataExportExecutionListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution list forbidden response has a 4xx status code
func (o *GetScanDataExportExecutionListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan data export execution list forbidden response has a 5xx status code
func (o *GetScanDataExportExecutionListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution list forbidden response a status code equal to that given
func (o *GetScanDataExportExecutionListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get scan data export execution list forbidden response
func (o *GetScanDataExportExecutionListForbidden) Code() int {
	return 403
}

func (o *GetScanDataExportExecutionListForbidden) Error() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListForbidden  %+v", 403, o.Payload)
}

func (o *GetScanDataExportExecutionListForbidden) String() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListForbidden  %+v", 403, o.Payload)
}

func (o *GetScanDataExportExecutionListForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionListNotFound creates a GetScanDataExportExecutionListNotFound with default headers values
func NewGetScanDataExportExecutionListNotFound() *GetScanDataExportExecutionListNotFound {
	return &GetScanDataExportExecutionListNotFound{}
}

/*
GetScanDataExportExecutionListNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetScanDataExportExecutionListNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution list not found response has a 2xx status code
func (o *GetScanDataExportExecutionListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution list not found response has a 3xx status code
func (o *GetScanDataExportExecutionListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution list not found response has a 4xx status code
func (o *GetScanDataExportExecutionListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan data export execution list not found response has a 5xx status code
func (o *GetScanDataExportExecutionListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution list not found response a status code equal to that given
func (o *GetScanDataExportExecutionListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get scan data export execution list not found response
func (o *GetScanDataExportExecutionListNotFound) Code() int {
	return 404
}

func (o *GetScanDataExportExecutionListNotFound) Error() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListNotFound  %+v", 404, o.Payload)
}

func (o *GetScanDataExportExecutionListNotFound) String() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListNotFound  %+v", 404, o.Payload)
}

func (o *GetScanDataExportExecutionListNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionListInternalServerError creates a GetScanDataExportExecutionListInternalServerError with default headers values
func NewGetScanDataExportExecutionListInternalServerError() *GetScanDataExportExecutionListInternalServerError {
	return &GetScanDataExportExecutionListInternalServerError{}
}

/*
GetScanDataExportExecutionListInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetScanDataExportExecutionListInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution list internal server error response has a 2xx status code
func (o *GetScanDataExportExecutionListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution list internal server error response has a 3xx status code
func (o *GetScanDataExportExecutionListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution list internal server error response has a 4xx status code
func (o *GetScanDataExportExecutionListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scan data export execution list internal server error response has a 5xx status code
func (o *GetScanDataExportExecutionListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scan data export execution list internal server error response a status code equal to that given
func (o *GetScanDataExportExecutionListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get scan data export execution list internal server error response
func (o *GetScanDataExportExecutionListInternalServerError) Code() int {
	return 500
}

func (o *GetScanDataExportExecutionListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanDataExportExecutionListInternalServerError) String() string {
	return fmt.Sprintf("[GET /export/cve/executions][%d] getScanDataExportExecutionListInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanDataExportExecutionListInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}