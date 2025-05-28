package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_hotel_reservation/types"
)

func HandleGetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(types.User{
		ID:        "",
		FirstName: "Jamal",
		LastName:  "Kurbanov",
	})
}

func HandleGetUser(ctx *fiber.Ctx) error {
	return ctx.JSON("James")
}
