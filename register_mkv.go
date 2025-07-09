package gomakemkv

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/DrWalrus1/gomakemkv/internal/streamReader"
)

const registerMkvKeySavedPrefix string = "Registration key saved"
const badKeyPrefix string = "Key not found or invalid"

// If the key used to register the instance of makemkv was invalid
var ErrBadKey = errors.New("key not found or invalid")

// This is a catch all errors not anticipated
var ErrUnexpectedRegistrationError = errors.New("unexpected error occurred while trying to register MakeMKV")

// Very simple parser for the makemkvcon reg command.
// Returns nil on success
// Returns [ErrBadKey] when key is invalid
// Returns [ErrUnexpectedRegistrationError] on unexpected response
func HandleRegisterMakeMkvEvents(reader io.Reader) error {
	c := streamReader.ReadStream(reader)
	for s := range c {
		s = strings.TrimSpace(s)
		switch {
		case strings.HasPrefix(s, badKeyPrefix):
			fmt.Println(s)
			return ErrBadKey
		case strings.HasPrefix(s, registerMkvKeySavedPrefix):
			fmt.Println(s)
			return nil
		default:
			fmt.Println(s)
		}
	}
	return ErrUnexpectedRegistrationError
}
