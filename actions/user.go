package actions

import (
	"transfr_backend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/uuid"
)

var db = make(map[uuid.UUID]models.User)

// UserResource as a helper method
type UserResource struct{}

// UserHandler declares the actions for the users in the DB
func (ur UserResource) UserHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Users!"}))
}

// CreateUser creates a user
func (ur UserResource) CreateUser(c buffalo.Context) error {
	user := &models.User{
		ID: uuid.Must(uuid.NewV4()),
	}
	db[user.ID] = *user
	return c.Render(201, r.JSON(user))
}
