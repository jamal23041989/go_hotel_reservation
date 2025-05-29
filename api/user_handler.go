package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jamal23041989/go_hotel_reservation/db"
	"github.com/jamal23041989/go_hotel_reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(types.User{
		FirstName: "Jamal",
		LastName:  "Kurbanov",
	})
}

func (h *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	var (
		id = ctx.Params("id")
		c  = context.Background()
	)
	user, err := h.userStore.GetUserByID(c, id)
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}
