package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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

// UserResponse is a model of the data coming back from the PUT requests
type UserResponse struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Name          string    `json:"name" db:"name"`
	Email         string    `json:"email" db:"email"`
	Password      string    `json:"password" db:"password"`
	WalletAddress string    `json:"wallet_address" db:"wallet_address"`
}

// CreateUser creates a user
func (ur UserResource) CreateUser(c buffalo.Context) error {
	// Establish a new user object;
	user := models.User{}
	var u UserResponse
	// Grab the Request coming in
	req := c.Request()

	// Decode the JSON
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	user.ID = uuid.Must(uuid.NewV4())
	user.Email = u.Email
	user.Name = u.Name
	user.Password = u.Password

	fmt.Println("Request Stuff", req)
	log.Println("Request Body", u.Email)
	tx.Create(&user)

	return c.Render(201, r.JSON(user))
}
