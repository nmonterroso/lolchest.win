package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nmonterroso/lolchest.win/riotapi/models"
)

// GetSummonerProfileReader is a Reader for the GetSummonerProfile structure.
type GetSummonerProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSummonerProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSummonerProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSummonerProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetSummonerProfileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSummonerProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetSummonerProfileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSummonerProfileOK creates a GetSummonerProfileOK with default headers values
func NewGetSummonerProfileOK() *GetSummonerProfileOK {
	return &GetSummonerProfileOK{}
}

/*GetSummonerProfileOK handles this case with default header values.

list of summoner profile data
*/
type GetSummonerProfileOK struct {
	Payload GetSummonerProfileOKBodyBody
}

func (o *GetSummonerProfileOK) Error() string {
	return fmt.Sprintf("[GET /api/lol/{region}/v1.4/summoner/by-name/{summonerNames}][%d] getSummonerProfileOK  %+v", 200, o.Payload)
}

func (o *GetSummonerProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSummonerProfileNotFound creates a GetSummonerProfileNotFound with default headers values
func NewGetSummonerProfileNotFound() *GetSummonerProfileNotFound {
	return &GetSummonerProfileNotFound{}
}

/*GetSummonerProfileNotFound handles this case with default header values.

summoner id or platform id not found
*/
type GetSummonerProfileNotFound struct {
}

func (o *GetSummonerProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/lol/{region}/v1.4/summoner/by-name/{summonerNames}][%d] getSummonerProfileNotFound ", 404)
}

func (o *GetSummonerProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSummonerProfileTooManyRequests creates a GetSummonerProfileTooManyRequests with default headers values
func NewGetSummonerProfileTooManyRequests() *GetSummonerProfileTooManyRequests {
	return &GetSummonerProfileTooManyRequests{}
}

/*GetSummonerProfileTooManyRequests handles this case with default header values.

rate limit exceeded
*/
type GetSummonerProfileTooManyRequests struct {
	/*the number of seconds to wait until retrying
	 */
	RetryAfter int32
}

func (o *GetSummonerProfileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/lol/{region}/v1.4/summoner/by-name/{summonerNames}][%d] getSummonerProfileTooManyRequests ", 429)
}

func (o *GetSummonerProfileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header retry-after
	retryAfter, err := swag.ConvertInt32(response.GetHeader("retry-after"))
	if err != nil {
		return errors.InvalidType("retry-after", "header", "int32", response.GetHeader("retry-after"))
	}
	o.RetryAfter = retryAfter

	return nil
}

// NewGetSummonerProfileInternalServerError creates a GetSummonerProfileInternalServerError with default headers values
func NewGetSummonerProfileInternalServerError() *GetSummonerProfileInternalServerError {
	return &GetSummonerProfileInternalServerError{}
}

/*GetSummonerProfileInternalServerError handles this case with default header values.

internal server error
*/
type GetSummonerProfileInternalServerError struct {
}

func (o *GetSummonerProfileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/lol/{region}/v1.4/summoner/by-name/{summonerNames}][%d] getSummonerProfileInternalServerError ", 500)
}

func (o *GetSummonerProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSummonerProfileServiceUnavailable creates a GetSummonerProfileServiceUnavailable with default headers values
func NewGetSummonerProfileServiceUnavailable() *GetSummonerProfileServiceUnavailable {
	return &GetSummonerProfileServiceUnavailable{}
}

/*GetSummonerProfileServiceUnavailable handles this case with default header values.

service unavailable
*/
type GetSummonerProfileServiceUnavailable struct {
}

func (o *GetSummonerProfileServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/lol/{region}/v1.4/summoner/by-name/{summonerNames}][%d] getSummonerProfileServiceUnavailable ", 503)
}

func (o *GetSummonerProfileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetSummonerProfileOKBodyBody get summoner profile o k body body

swagger:model GetSummonerProfileOKBodyBody
*/
type GetSummonerProfileOKBodyBody map[string]models.SummonerDto

// Validate validates this get summoner profile o k body body
func (o GetSummonerProfileOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("getSummonerProfileOK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
