package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullPath(t *testing.T) {
	var testdata = []struct {
		workspace string
		dir       string
		name      string
		expect    string
	}{
		{
			workspace: "/test1/",
			dir:       "/test2/",
			name:      "test3.json",
			expect:    "/test1/test2/test3.json",
		},
		{
			workspace: "/test1/",
			dir:       "/test2",
			name:      "test3.json",
			expect:    "/test1/test2/test3.json",
		},
		{
			workspace: "/test1",
			dir:       "/test2",
			name:      "test3.json",
			expect:    "/test1/test2/test3.json",
		},
		{
			workspace: "/test1/",
			dir:       "/test2/",
			name:      "/test3.json",
			expect:    "/test1/test2/test3.json",
		},
	}

	for _, v := range testdata {
		value := FullPath(v.workspace, v.dir, v.name)
		assert.Equal(t, v.expect, value)
	}
}
