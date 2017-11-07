package chapter1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// ConvertHexToBase64 ...
// takes a hex string and converts it to base64 string
func ConvertHexToBase64(h string) string {
	decoded, err := hex.DecodeString(h)

	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(decoded)
}
