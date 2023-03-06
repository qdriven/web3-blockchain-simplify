package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()
	fmt.Println(config)
	assert.Equal(t, config.Ws.WorkingDir, ".")
}
