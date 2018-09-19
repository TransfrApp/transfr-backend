package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Coupon declares the model structure for coupons
type Coupon struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	Name               string    `json:"name" db:"name"`
	Code               string    `json:"code" db:"code"`
	DiscountPercentage int       `json:"discount_percentage" db:"discount_percentage"`
	DiscountSetValue   float64   `json:"discount_set_value" db:"discount_set_value"`
	User               User      `belongs_to:"users" db:"u_id"`
}

// String is not required by pop and may be deleted
func (c Coupon) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Coupons is not required by pop and may be deleted
type Coupons []Coupon

// String is not required by pop and may be deleted
func (c Coupons) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Coupon) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.Code, Name: "Code"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Coupon) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Coupon) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
