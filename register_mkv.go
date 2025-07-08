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

var ErrBadKey = errors.New("key not found or invalid")
var ErrUnexpectedRegistrationError = errors.New("unexpected error occurred while trying to register MakeMKV")

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
