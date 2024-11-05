package api

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
	"wav-to-flac-converter/audio"
)

type ErrorMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// sendErrorMessage sends an error message to the client in JSON format
func sendErrorMessage(c *websocket.Conn, message string) {
	errorMsg := ErrorMessage{
		Type:    "error",
		Message: message,
	}
	jsonData, _ := json.Marshal(errorMsg)
	c.WriteMessage(websocket.TextMessage, jsonData)
}

func WebSocketHandler(c *websocket.Conn) {
	defer c.Close()

	for {
		// Read WAV data from WebSocket
		_, wavData, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			sendErrorMessage(c, "Failed to read WAV data from WebSocket")
			break
		}

		// Convert WAV to FLAC
		flacData, err := audio.ConvertWAVToFLAC(wavData)
		if err != nil {
			log.Println("Error converting WAV to FLAC:", err)
			sendErrorMessage(c, "Failed to convert WAV data to FLAC format")
			break
		}

		// Stream FLAC data back to WebSocket
		if err := c.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
			log.Println("Error sending FLAC data:", err)
			sendErrorMessage(c, "Failed to send FLAC data back to client")
			break
		}
	}
}
