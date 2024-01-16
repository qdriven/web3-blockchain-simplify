package misc

import (
	"encoding/base64"
	"log"
	"os"
)

func ConvertToBase64(imagePath string) string {
	data, err := os.ReadFile(imagePath)
	if err != nil {
		log.Fatal("image read failed,", err)
	}
	result := base64.StdEncoding.EncodeToString(data)
	return result
}
