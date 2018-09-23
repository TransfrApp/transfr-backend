package actions

import (
	"encoding/json"
	"fmt"
	"time"
	"transfr_backend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
)

// var db = make(map[uuid.UUID]models.User)

// UserResource as a helper method
type UserResource struct{}

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

// GetAllUsers declares the actions for the users in the DB
func (ur UserResource) GetAllUsers(c buffalo.Context) error {
	db, err := pop.Connect("development")
	if err != nil {
		fmt.Println("Connection Issue", err)
	}
	users := []models.User{}
	userErr := db.All(&users)
	if userErr != nil {
		fmt.Println("Issue with find all user", userErr)
	}
	return c.Render(200, r.JSON(users))
}

// GetUser gets back a single user with the users id
// This route is declared at "/users/{id}"
func (ur UserResource) GetUser(c buffalo.Context) error {
	// Connect to the DB
	db, err := pop.Connect("development")
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
	}
	// connect to user model
	user := models.User{}

	//pull back param from URL
	// get id and format to uuid
	id, userErr := uuid.FromString(c.Param("id"))
	if userErr != nil {
		return c.Render(200, r.JSON(map[string]string{"error": "There was an error with the params you passed in"}))
	}

	// Find the user in the DB
	findErr := db.Find(&user, id)
	if findErr != nil {
		fmt.Println("Error finding user", findErr)
	}
	return c.Render(200, r.JSON(user))
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
	user.WalletAddress = u.WalletAddress

	tx.Create(&user)

	return c.Render(201, r.JSON(user))
}

// UpdateUser takes the user and updates it's current values
// Routed to "/users/{id}"
func UpdateUser(c buffalo.Context) error {
	// establish connction to the db
	db, err := pop.Connect("development")
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Connection to user model
	user := models.User{}
	var u UserResponse

	// Grab the users id
	id, userErr := uuid.FromString(c.Param("id"))
	if userErr != nil {
		return c.Render(200, r.JSON(map[string]string{"error": "There was an error with the params you passed in"}))
	}
	fmt.Println("ID from params", id)
	// Grab info from the API call
	req := c.Request()
	decoder := json.NewDecoder(req.Body)
	decodeErr := decoder.Decode(&u)
	if decodeErr != nil {
		panic(decodeErr)
	}
	fmt.Println("Request Body", u)
	//find user in Db
	findUserErr := db.Find(&user, id)
	if findUserErr != nil {
		fmt.Println("Error finding users", findUserErr)
	}
	fmt.Println("User getting update", user)
	fmt.Println("Testing the u.Email thing", u.Email)
	//Update the users info
	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.Password
	user.WalletAddress = u.WalletAddress
	updateUserErr := tx.Update(&user)
	if updateUserErr != nil {
		fmt.Println("Update user error", updateUserErr)
	}
	return c.Render(200, r.JSON(user))
}

// DeleteUser deletes the user
func (ur UserResource) DeleteUser(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(user); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(map[string]string{"success": "Hell yeah, that mother fucker just got wrecked"}))
}
