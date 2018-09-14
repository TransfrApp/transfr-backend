package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Inventory struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	Quantity  int       `json:"quantity" db:"quantity"`
	MetaTags  string    `json:"meta_tags" db:"meta_tags"`
}

// String is not required by pop and may be deleted
func (i Inventory) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Inventories is not required by pop and may be deleted
type Inventories []Inventory

// String is not required by pop and may be deleted
func (i Inventories) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Inventory) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Name, Name: "Name"},
		&validators.StringIsPresent{Field: i.MetaTags, Name: "MetaTags"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Inventory) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Inventory) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
