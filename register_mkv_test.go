package gomakemkv_test

import (
	"strings"
	"testing"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/go-playground/assert/v2"
)

func TestRegisterMakeMkvEventHandlerInvalidKey(t *testing.T) {
	input := "Key not found or invalid"
	responseCode := gomakemkv.HandleRegisterMakeMkvEvents(strings.NewReader(input))

	assert.Equal(t, gomakemkv.ErrBadKey, responseCode)
}

func TestRegisterMakeMkvEventHandlerValidKey(t *testing.T) {
	input := "Registration key saved -- some extra dummy data"
	responseCode := gomakemkv.HandleRegisterMakeMkvEvents(strings.NewReader(input))

	assert.Equal(t, nil, responseCode)
}

func TestRegisterMakeMkvEventHandlerValidKeyKeyAlreadyExists(t *testing.T) {
	input := `Current Key already exists: asdljasdlkjalkdjlk
Registration key saved -- some extra dummy data`
	responseCode := gomakemkv.HandleRegisterMakeMkvEvents(strings.NewReader(input))

	assert.Equal(t, nil, responseCode)
}
