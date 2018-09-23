package actions

import (
	"github.com/gobuffalo/buffalo"
)

// GetInventory is a default handler to serve up
// a home page.
func GetInventory(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Get Inventory!"}))
}
