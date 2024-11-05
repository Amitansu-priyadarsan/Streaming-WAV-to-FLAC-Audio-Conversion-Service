# 🎵 WAV to FLAC Audio Conversion Service

## Project Overview

The WAV to FLAC Audio Conversion Service is a robust backend service developed in Go that allows real-time conversion of WAV audio streams to FLAC format. Utilizing WebSockets, the service efficiently streams converted audio data back to the client, ensuring minimal latency and high fidelity.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Load Testing with k6](#load-testing-with-k6)
- [Performance Considerations](#performance-considerations)
- [Conclusion](#conclusion)

## Features

- 🎶 **Real-time Audio Conversion**: Streams WAV audio data and converts it to FLAC format on-the-fly.
- ⚡ **WebSocket Communication**: Uses WebSocket for real-time data transmission.
- ❗ **Error Handling**: Enhanced error handling during audio processing and streaming.
- ✅ **Comprehensive Testing**: Includes unit and integration tests to validate functionality.
- 📊 **Load Testing**: Supports performance testing using k6.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker (optional for containerization)
- WebSocket library (`github.com/gorilla/websocket`)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/wav-to-flac-converter.git
   cd wav-to-flac-converter

2.Install dependencies:
```bash
   go mod tidy
   ```
3.Run the service:
```bash
   go run main.go

   ```

4.The server will start listening on localhost:3000. Ensure you have a WAV file to test the service.

## Client
To demonstrate functionality, use the provided client.go to send a WAV file and receive the converted FLAC file.

1.Place a WAV file named sample.wav in the same directory as client.go.

2.Run the client:
```bash
go run client.go
```
3.The output FLAC file will be saved as output.flac.

## API Endpoints

### WebSocket Endpoint
- **URL**: `/ws`
- **Method**: `POST`
- **Description**: Accepts a WAV audio stream and returns the converted FLAC audio stream.

#### Request
- **Content-Type**: `application/octet-stream`
- **Body**: WAV audio data in binary format.

###Response
- Content-Type: audio/flac
- Body: FLAC audio data

## Testing
Unit Testing
To run unit tests for the WAV to FLAC conversion logic, execute the following command:
```bash
go test ./tests/...
```
## Integration Testing
To validate WebSocket connections and streaming capabilities, run:
```bash
go test ./tests/integration_test.go
```
## Load Testing with k6
To perform load testing on the WebSocket endpoint, use the provided loadtest.js script.

Install k6 if you haven't already:
```bash
brew install k6
```
Run the load test:
```bash
k6 run loadtest.js
```
Output will look like -
```bash
data_received.........: 0 B  0 B/s
     data_sent.............: 0 B  0 B/s
     iteration_duration....: avg=1s       min=1s      med=1s       max=1s     p(90)=1s     p(95)=1s    
     iterations............: 3000 99.793292/s
     vus...................: 100  min=100     max=100
     vus_max...............: 100  min=100     max=100
     ws_connecting.........: avg=519.49µs min=52.29µs med=245.22µs max=6.87ms p(90)=1.23ms p(95)=1.77ms
     ws_session_duration...: avg=529.14µs min=56.95µs med=258.08µs max=6.87ms p(90)=1.24ms p(95)=1.77ms
     ws_sessions...........: 3000 99.793292/s
```


## Performance Considerations
- ⏱️ Low-Latency Processing: The service is optimized for minimal delay during audio processing.
- 📈 Scalability: Can be deployed in a containerized environment (e.g., Kubernetes) to handle high traffic loads.
- 🚨 Error Handling: The system is designed to gracefully handle errors during streaming and conversion processes.

## Conclusion
The WAV to FLAC Audio Conversion Service is a scalable and efficient backend solution for real-time audio processing. Comprehensive documentation, testing strategies, and performance optimizations ensure a reliable service that meets the needs of audio streaming applications.


