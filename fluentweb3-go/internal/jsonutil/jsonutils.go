package jsonutil

import (
	"encoding/json"
	"log"
	"os"
)

func ToStruct(jsonstr string, any any) {

	err := json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}

func ToStructureString(any any) string {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}

func ToStructFromFile(jsonFile string, any any) {
	jsonstr, err := os.ReadFile(jsonFile)
	err = json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		log.Fatal("convert json str failed", any)
	}
}
