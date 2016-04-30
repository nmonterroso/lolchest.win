package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/nmonterroso/lolchest.win/models"
)

/*GetAllChampDataOK all champion static data

swagger:response getAllChampDataOK
*/
type GetAllChampDataOK struct {

	// In: body
	Payload []*models.ChampionData `json:"body,omitempty"`
}

// NewGetAllChampDataOK creates GetAllChampDataOK with default headers values
func NewGetAllChampDataOK() *GetAllChampDataOK {
	return &GetAllChampDataOK{}
}

// WithPayload adds the payload to the get all champ data o k response
func (o *GetAllChampDataOK) WithPayload(payload []*models.ChampionData) *GetAllChampDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all champ data o k response
func (o *GetAllChampDataOK) SetPayload(payload []*models.ChampionData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChampDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetAllChampDataInternalServerError unexpected error

swagger:response getAllChampDataInternalServerError
*/
type GetAllChampDataInternalServerError struct {
}

// NewGetAllChampDataInternalServerError creates GetAllChampDataInternalServerError with default headers values
func NewGetAllChampDataInternalServerError() *GetAllChampDataInternalServerError {
	return &GetAllChampDataInternalServerError{}
}

// WriteResponse to the client
func (o *GetAllChampDataInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}

/*GetAllChampDataBadGateway invalid response from riot api

swagger:response getAllChampDataBadGateway
*/
type GetAllChampDataBadGateway struct {
}

// NewGetAllChampDataBadGateway creates GetAllChampDataBadGateway with default headers values
func NewGetAllChampDataBadGateway() *GetAllChampDataBadGateway {
	return &GetAllChampDataBadGateway{}
}

// WriteResponse to the client
func (o *GetAllChampDataBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
}
