package internal

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetVersion(t *testing.T) {
	got := GetVersion()
	want := fmt.Sprintf(`Version:%s Revision: %s OS:%s Arch:%s`, MAJOR_VERSION, Revision, runtime.GOOS, runtime.GOARCH)

	if want != got {
		t.Fatalf("unexpected version info. want: %s, got: %s", want, got)
	}
}
