package ollamahdl

import (
	"fmt"
	"ollamawithgo/internal/service/ollamasvc/llamasvc"

	"github.com/gofiber/fiber/v2"
)

type LlamaHandler struct {
	llama llamasvc.Llama3_2Service
}

func NewLlamaHandler(llama llamasvc.Llama3_2Service) *LlamaHandler {
	return &LlamaHandler{
		llama: llama,
	}
}

func (h *LlamaHandler) Init(root fiber.Router) {
	root.Post("/generateCompletionFull", h.GenerateCompletion)
	root.Post("/generateCompletionStream", h.GenerateCompletionStream)
}

func (h *LlamaHandler) GenerateCompletion(c *fiber.Ctx) error {
	var req llamasvc.GenerateACompletionServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.llama.GenerateACompletionFull(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"response": resp})
}

func (h *LlamaHandler) GenerateCompletionStream(c *fiber.Ctx) error {
	var req llamasvc.GenerateACompletionServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Get the streaming response channel from the llama model
	respChan, err := h.llama.GenerateACompletionStream(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Set the Content-Type to JSON
	c.Set("Content-Type", "application/json")

	// Use SendStream to stream the response to the client
	for resp := range respChan {
		// joinedResponse := strings.Join(resp, "")
		if _, err := c.Context().Write([]byte(resp[0])); err != nil {
			fmt.Println("Error while streaming response:", err)
			return err
		}
	}

	// You can optionally handle additional status or logic here
	return nil
}
