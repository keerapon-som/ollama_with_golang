package handler

import (
	"ollamawithgo/internal/service/usersvc"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService usersvc.UserService
}

func NewUserHandler(userService usersvc.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Init(root fiber.Router) {
	root.Get("/users", h.GetAllUsers)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	GetAllUsers, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"users": GetAllUsers})
}
