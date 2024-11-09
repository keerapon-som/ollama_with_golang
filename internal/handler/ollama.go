package handler

import (
	"ollamawithgo/internal/service/ollamasvc"

	"github.com/gofiber/fiber/v2"
)

type OllamaHandler struct {
	ollamaService ollamasvc.OllamaService
}

func NewOllamaHandler(OllamaService ollamasvc.OllamaService) *OllamaHandler {
	return &OllamaHandler{
		ollamaService: OllamaService,
	}
}

func (h *OllamaHandler) Init(root fiber.Router) {
	root.Get("/models", h.GetAllModels)
}

func (h *OllamaHandler) GetAllModels(c *fiber.Ctx) error {
	AllModels, err := h.ollamaService.GetAllModelsName()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"models": AllModels})
}
