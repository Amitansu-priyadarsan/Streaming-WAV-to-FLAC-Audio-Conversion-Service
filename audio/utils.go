// audio/utils.go
package audio

import "log"

// ErrorLogger logs errors if they occur
func ErrorLogger(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
