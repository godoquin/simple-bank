package util

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
