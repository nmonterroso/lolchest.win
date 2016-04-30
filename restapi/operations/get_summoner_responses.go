package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nmonterroso/lolchest.win/models"
)

/*GetSummonerOK summoner data

swagger:response getSummonerOK
*/
type GetSummonerOK struct {

	// In: body
	Payload *models.Summoner `json:"body,omitempty"`
}

// NewGetSummonerOK creates GetSummonerOK with default headers values
func NewGetSummonerOK() *GetSummonerOK {
	return &GetSummonerOK{}
}

// WithPayload adds the payload to the get summoner o k response
func (o *GetSummonerOK) WithPayload(payload *models.Summoner) *GetSummonerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get summoner o k response
func (o *GetSummonerOK) SetPayload(payload *models.Summoner) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSummonerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSummonerInternalServerError unexpected error

swagger:response getSummonerInternalServerError
*/
type GetSummonerInternalServerError struct {
}

// NewGetSummonerInternalServerError creates GetSummonerInternalServerError with default headers values
func NewGetSummonerInternalServerError() *GetSummonerInternalServerError {
	return &GetSummonerInternalServerError{}
}

// WriteResponse to the client
func (o *GetSummonerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}

/*GetSummonerBadGateway invalid response from riot api

swagger:response getSummonerBadGateway
*/
type GetSummonerBadGateway struct {
}

// NewGetSummonerBadGateway creates GetSummonerBadGateway with default headers values
func NewGetSummonerBadGateway() *GetSummonerBadGateway {
	return &GetSummonerBadGateway{}
}

// WriteResponse to the client
func (o *GetSummonerBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
}
