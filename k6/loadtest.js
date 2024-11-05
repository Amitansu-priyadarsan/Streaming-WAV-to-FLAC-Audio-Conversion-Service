import { check, sleep } from "k6";
import ws from "k6/ws";

// Directly load sample WAV data as binary
const wavData = open("sample.wav", "b"); // "b" specifies binary mode

export const options = {
    vus: 100,           // Number of virtual users
    duration: "30s",    // Duration of the test
};

export default function () {
    const url = "ws://localhost:3000/ws";

    ws.connect(url, null, function (socket) {
        socket.on("open", () => {
            console.log("Connected to WebSocket");
            socket.sendBinary(wavData); // Send the WAV data as binary
        });

        socket.on("message", (message) => {
            // Check if received message contains data (FLAC data)
            check(message, {
                "received FLAC data": (msg) => msg.length > 0,
            });
            socket.close(); // Close connection after receiving FLAC data
        });

        socket.on("error", (e) => {
            console.error("WebSocket error:", e.error());
        });

        socket.on("close", () => {
            console.log("Disconnected from WebSocket");
        });
    });

    sleep(1); // Wait before the next virtual user connects
}
