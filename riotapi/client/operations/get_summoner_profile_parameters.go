package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSummonerProfileParams creates a new GetSummonerProfileParams object
// with the default values initialized.
func NewGetSummonerProfileParams() *GetSummonerProfileParams {
	var (
		regionDefault string = string("na")
	)
	return &GetSummonerProfileParams{
		Region: regionDefault,
	}
}

/*GetSummonerProfileParams contains all the parameters to send to the API endpoint
for the get summoner profile operation typically these are written to a http.Request
*/
type GetSummonerProfileParams struct {

	/*Region*/
	Region string
	/*SummonerNames*/
	SummonerNames string
}

// WithRegion adds the region to the get summoner profile params
func (o *GetSummonerProfileParams) WithRegion(Region string) *GetSummonerProfileParams {
	o.Region = Region
	return o
}

// WithSummonerNames adds the summonerNames to the get summoner profile params
func (o *GetSummonerProfileParams) WithSummonerNames(SummonerNames string) *GetSummonerProfileParams {
	o.SummonerNames = SummonerNames
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSummonerProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	// path param summonerNames
	if err := r.SetPathParam("summonerNames", o.SummonerNames); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
