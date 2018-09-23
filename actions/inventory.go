package actions

import (
	"github.com/gobuffalo/buffalo"
)

//TransactionResource is a support
type TransactionResource struct{}

// GetInventory is a default handler to serve up
// a home page.
func GetInventory(c buffalo.Context) error {
	db, err := pop.Context("development")
	if err != nil {
		fmt.Println("Issue connection to DB", err)
	}
	//Connect to models
	inventory := models.Inventory{}
	// Find all items in inventory
	findErr := db.All(&inventory)
	if findErr != nil {
		fmt.Println("Issue with find items", findErr)
	}
	return c.Render(200, r.JSON(200, r.JSON(inventory))
}
