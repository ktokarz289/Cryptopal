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

// XOR ...
// takes two strings and converts to XOR string
func XOR(s1, s2 []byte) string {
	prod := make([]byte, len(s1))
	for i := 0; i < len(prod); i++ {
		prod[i] = s1[i] ^ s2[i]
	}

	return hex.EncodeToString(prod)
}
