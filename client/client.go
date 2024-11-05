package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/url"
    "os"

    "github.com/gorilla/websocket"
)

func main() {
    // Open the WAV file located in the same directory as this script
    wavFilePath := "sample.wav" // Updated this path to your WAV file
    wavFile, err := os.Open(wavFilePath)
    if err != nil {
        log.Fatal("Error opening WAV file:", err)
    }
    defer wavFile.Close()

    // Read the WAV file into a byte slice
    wavData, err := ioutil.ReadAll(wavFile)
    if err != nil {
        log.Fatal("Error reading WAV file:", err)
    }

    // Connect to the WebSocket server
    u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
    conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("Error connecting to WebSocket:", err)
    }
    defer conn.Close()

    // Send the WAV file data to the server
    err = conn.WriteMessage(websocket.BinaryMessage, wavData)
    if err != nil {
        log.Fatal("Error sending WAV data:", err)
    }

    // Read the response (the FLAC data)
    _, flacData, err := conn.ReadMessage()
    if err != nil {
        log.Fatal("Error reading response:", err)
    }

    // Write the FLAC data to a file
    flacFilePath := "output.flac"
    err = ioutil.WriteFile(flacFilePath, flacData, 0644)
    if err != nil {
        log.Fatal("Error writing FLAC file:", err)
    }

    fmt.Println("Conversion successful! FLAC file saved as:", flacFilePath)
}
