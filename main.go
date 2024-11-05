package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"wav-to-flac-converter/api"
)

func main() {
	app := fiber.New()

	// WebSocket route to handle WAV to FLAC streaming
	app.Get("/ws", websocket.New(api.WebSocketHandler))

	log.Println("Starting server on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
