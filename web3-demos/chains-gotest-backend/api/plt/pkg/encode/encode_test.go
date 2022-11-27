package encode

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestDurationText(t *testing.T) {
	var (
		input  = []byte("10s")
		output = time.Second * 10
		d      Duration
	)
	if err := d.UnmarshalText(input); err != nil {
		t.FailNow()
	}
	if int64(output) != int64(d) {
		t.FailNow()
	}

	enc, err := d.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, enc)
}
