package handler

import (
	"ollamawithgo/internal/ollama"
	"ollamawithgo/internal/service/ollamasvc"
	"ollamawithgo/internal/web"

	"github.com/gofiber/fiber/v2"
)

func CreateHandler() *fiber.App {
	app := fiber.New()

	// Enable CORS for all routes
	// app.Use(cors.New())
	// Public routes
	public := app.Group("/api")
	registerPublicHandlers(public)

	// private := app.Group("/api/private", middleware.AuthRequired)
	private := app.Group("/api/private")
	registerPrivateHandlers(private)
	return app
}

func registerPublicHandlers(app fiber.Router) {

	handler := web.HandlerRegistrator{}
	handler.Register(
		NewOllamaHandler(ollamasvc.NewOllamaService(ollama.NewOllama())),
	)
	handler.Init(app)
}

func registerPrivateHandlers(app fiber.Router) {
	handler := web.HandlerRegistrator{}
	handler.Register()
	handler.Init(app)
}
