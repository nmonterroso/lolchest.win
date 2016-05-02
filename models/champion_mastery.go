package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChampionMastery champion mastery

swagger:model ChampionMastery
*/
type ChampionMastery struct {

	/* champ icon Url

	Required: true
	*/
	ChampIconURL *string `json:"champIconUrl"`

	/* champ Id

	Required: true
	*/
	ChampID *int64 `json:"champId"`

	/* champ name

	Required: true
	*/
	ChampName *string `json:"champName"`

	/* chest is available

	Required: true
	*/
	ChestIsAvailable *bool `json:"chestIsAvailable"`

	/* highest grade
	 */
	HighestGrade string `json:"highestGrade,omitempty"`
}

// Validate validates this champion mastery
func (m *ChampionMastery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChampIconURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChampID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChampName(formats); err != nil {
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

func (m *ChampionMastery) validateChampIconURL(formats strfmt.Registry) error {

	if err := validate.Required("champIconUrl", "body", m.ChampIconURL); err != nil {
		return err
	}

	return nil
}

func (m *ChampionMastery) validateChampID(formats strfmt.Registry) error {

	if err := validate.Required("champId", "body", m.ChampID); err != nil {
		return err
	}

	return nil
}

func (m *ChampionMastery) validateChampName(formats strfmt.Registry) error {

	if err := validate.Required("champName", "body", m.ChampName); err != nil {
		return err
	}

	return nil
}

func (m *ChampionMastery) validateChestIsAvailable(formats strfmt.Registry) error {

	if err := validate.Required("chestIsAvailable", "body", m.ChestIsAvailable); err != nil {
		return err
	}

	return nil
}