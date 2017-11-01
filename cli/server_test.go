package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_runServer_Success(t *testing.T) {
	err := runServer(nil, nil)

	assert.NoError(t, err, "An error occured running the server CLI.")
}
