package image

import (
	"fmt"
	"testing"
)

func TestImageToBase64(t *testing.T) {

	result := ConvertToBase64("./datalayer01.webp")
	fmt.Println(result)
}
