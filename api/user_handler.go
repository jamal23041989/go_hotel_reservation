package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_hotel_reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Last",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("")
}
