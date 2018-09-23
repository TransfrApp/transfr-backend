package actions

import (
	"transfr_backend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

var db = make(map[uuid.UUID]models.User)

// UserResource as a helper method
type UserResource struct{}

// UserHandler declares the actions for the users in the DB
func (ur UserResource) UserHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}

// CreateUser creates a user
func (ur UserResource) CreateUser(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	user := models.User{}

	user.ID = uuid.Must(uuid.NewV4())
	user.Email = "Test Email"
	user.Name = "Test User"
	user.Password = "changeme"

	tx.Create(&user)

	return c.Render(201, r.JSON(user))
}
