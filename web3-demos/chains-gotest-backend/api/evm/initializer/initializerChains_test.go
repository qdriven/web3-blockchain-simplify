package initializer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBasePath(t *testing.T) {
	basePath :=GetProjectPath()
	assert.Equal(t, "/Users/Patrick/workspace/products/focus/go-chains/",basePath)
}

func TestSetup(t *testing.T) {
	Setup4Test()
}