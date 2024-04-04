package handlers

import "github.com/gofiber/fiber/v2"

type HTTPHandlers struct {
	app *fiber.App
}

func New(app *fiber.App) *HTTPHandlers {
	h := &HTTPHandlers{}

	h.app = app

	return h
}

func (h *HTTPHandlers) Init() {
	// Define a route handler
	// Subscribe to a queue
	h.app.Get("/subscribe", h.subscribe)

	// Get Events
	h.app.Get("/snitch", h.snitch)
}

func (h *HTTPHandlers) subscribe(c *fiber.Ctx) error {
	return c.SendString("subscribe")
}

func (h *HTTPHandlers) snitch(c *fiber.Ctx) error {
	return c.SendString("notified")
}
