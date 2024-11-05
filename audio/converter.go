package audio

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// ConvertWAVToFLAC converts WAV data to FLAC format using an external `flac` command-line tool.
func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
	// Write the WAV data to a temporary file
	tempWavFile, err := ioutil.TempFile("", "temp-*.wav")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp WAV file: %w", err)
	}
	defer tempWavFile.Close()
	defer os.Remove(tempWavFile.Name()) // Clean up temp file after function exits

	if _, err := tempWavFile.Write(wavData); err != nil {
		return nil, fmt.Errorf("failed to write to temp WAV file: %w", err)
	}

	// Prepare the FLAC output file
	tempFlacFile := tempWavFile.Name() + ".flac"

	// Run the `flac` command to convert WAV to FLAC
	cmd := exec.Command("flac", "-f", tempWavFile.Name(), "-o", tempFlacFile)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to convert WAV to FLAC: %w", err)
	}
	defer os.Remove(tempFlacFile) // Clean up temp FLAC file after function exits

	// Read the FLAC data from the output file
	flacData, err := ioutil.ReadFile(tempFlacFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read FLAC file: %w", err)
	}

	return flacData, nil
}
