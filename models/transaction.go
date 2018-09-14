package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Transaction struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	To        string    `json:"to" db:"to"`
	From      string    `json:"from" db:"from"`
	Amount    integer   `json:"amount" db:"amount"`
	Items     string    `json:"items" db:"items"`
}

// String is not required by pop and may be deleted
func (t Transaction) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Transactions is not required by pop and may be deleted
type Transactions []Transaction

// String is not required by pop and may be deleted
func (t Transactions) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Transaction) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.To, Name: "To"},
		&validators.StringIsPresent{Field: t.From, Name: "From"},
		&validators.StringIsPresent{Field: t.Items, Name: "Items"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Transaction) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Transaction) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
