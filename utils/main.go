package utils

import (
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
