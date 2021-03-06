package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetChampionDataParams creates a new GetChampionDataParams object
// with the default values initialized.
func NewGetChampionDataParams() *GetChampionDataParams {
	var (
		champDataDefault string = string("image")
		regionDefault    string = string("na")
	)
	return &GetChampionDataParams{
		ChampData: &champDataDefault,
		Region:    regionDefault,
	}
}

/*GetChampionDataParams contains all the parameters to send to the API endpoint
for the get champion data operation typically these are written to a http.Request
*/
type GetChampionDataParams struct {

	/*ChampData*/
	ChampData *string
	/*Region*/
	Region string
}

// WithChampData adds the champData to the get champion data params
func (o *GetChampionDataParams) WithChampData(ChampData *string) *GetChampionDataParams {
	o.ChampData = ChampData
	return o
}

// WithRegion adds the region to the get champion data params
func (o *GetChampionDataParams) WithRegion(Region string) *GetChampionDataParams {
	o.Region = Region
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetChampionDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.ChampData != nil {

		// query param champData
		var qrChampData string
		if o.ChampData != nil {
			qrChampData = *o.ChampData
		}
		qChampData := qrChampData
		if qChampData != "" {
			if err := r.SetQueryParam("champData", qChampData); err != nil {
				return err
			}
		}

	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
