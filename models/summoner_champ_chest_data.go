package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*SummonerChampChestData summoner champ chest data

swagger:model SummonerChampChestData
*/
type SummonerChampChestData struct {

	/* champ Id

	Required: true
	*/
	ChampID *int64 `json:"champId"`

	/* chest is available

	Required: true
	*/
	ChestIsAvailable *bool `json:"chestIsAvailable"`
}

// Validate validates this summoner champ chest data
func (m *SummonerChampChestData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChampID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChestIsAvailable(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummonerChampChestData) validateChampID(formats strfmt.Registry) error {

	if err := validate.Required("champId", "body", m.ChampID); err != nil {
		return err
	}

	return nil
}

func (m *SummonerChampChestData) validateChestIsAvailable(formats strfmt.Registry) error {

	if err := validate.Required("chestIsAvailable", "body", m.ChestIsAvailable); err != nil {
		return err
	}

	return nil
}
