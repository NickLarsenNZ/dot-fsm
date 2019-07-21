package utils

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

func TextToIdentifier(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToUpper(reg.ReplaceAllString(text, "_"))
}

func JsonDump(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		return ""
	}
	return string(bytes)
}
