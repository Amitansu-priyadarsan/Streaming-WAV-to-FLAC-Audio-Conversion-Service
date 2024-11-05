package tests

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2" // Fiber's WebSocket package
	"wav-to-flac-converter/api"
	"github.com/stretchr/testify/assert"
	gorillaWebsocket "github.com/gorilla/websocket" // Alias Gorilla's WebSocket package
)

func TestWebSocketHandler(t *testing.T) {
	// Set up Fiber app and WebSocket route
	app := fiber.New()
	app.Get("/ws", websocket.New(api.WebSocketHandler))

	// Start a Fiber server in a goroutine
	go func() {
		if err := app.Listen(":8080"); err != nil {
			t.Fatalf("Failed to start Fiber server: %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// Convert server URL to WebSocket protocol
	url := "ws://localhost:8080/ws"

	// Connect to the WebSocket server using Gorilla's websocket package
	conn, _, err := gorillaWebsocket.DefaultDialer.Dial(url, nil) // Use the aliased package
	assert.NoError(t, err, "WebSocket connection should succeed")
	defer conn.Close()

	// Load sample WAV data for testing
	wavData, err := ioutil.ReadFile("sample.wav") // Make sure sample.wav exists in the project root or specify the correct path
	assert.NoError(t, err, "Sample WAV file should be available for testing")

	// Send WAV data to WebSocket
	err = conn.WriteMessage(gorillaWebsocket.BinaryMessage, wavData) // Use the aliased package
	assert.NoError(t, err, "Writing message should succeed")

	// Wait and receive response
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, flacData, err := conn.ReadMessage()

	// Assertions
	assert.NoError(t, err, "Reading message should succeed")
	assert.NotNil(t, flacData, "FLAC data should not be nil")
	assert.Greater(t, len(flacData), 0, "FLAC data should not be empty")
}


