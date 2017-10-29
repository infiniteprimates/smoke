package cli

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestServer_runInit_Success(t *testing.T) {
	err := runInit(nil, nil);

	assert.NoError(t, err, "An error occured running the init CLI.")
}
