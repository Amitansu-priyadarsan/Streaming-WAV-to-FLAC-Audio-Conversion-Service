package tests

import (
	"io/ioutil"
	"testing"
	"wav-to-flac-converter/audio"

	"github.com/stretchr/testify/assert"
)

func TestConvertWAVToFLAC(t *testing.T) {
	// Load a small sample WAV file for testing
	wavData, err := ioutil.ReadFile("sample.wav") // Make sure sample.wav exists in the project root or specify the correct path
	if err != nil {
		t.Fatalf("Failed to read sample WAV file: %v", err)
	}

	// Perform the conversion
	flacData, err := audio.ConvertWAVToFLAC(wavData)

	// Assertions
	assert.NoError(t, err, "Conversion should not return an error")
	assert.NotNil(t, flacData, "FLAC data should not be nil")
	assert.Greater(t, len(flacData), 0, "FLAC data should not be empty")
}
