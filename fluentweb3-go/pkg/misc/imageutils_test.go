package misc

import (
	"fmt"
	"testing"
)

func TestImageToBase64(t *testing.T) {

	result := ConvertToBase64("./img.png")
	fmt.Println(result)
}
