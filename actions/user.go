package actions

import "github.com/gobuffalo/buffalo"

// UserHandler declares the actions for the users in the DB
func UserHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Users!"}))
}
