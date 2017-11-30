package chapter1_test

import (
	"encoding/hex"
	"testing"

	"github.com/ktokarz289/Cryptopal/Chapters/chapter1"
)

func TestHexToBase64(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	value := chapter1.ConvertHexToBase64(hex)

	if value != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Error()
	}
}

func TestXor(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	value, _ := hex.DecodeString(hex1)
	value2, _ := hex.DecodeString(hex2)

	val := chapter1.XOR(value, value2)

	if val != "746865206b696420646f6e277420706c6179" {
		t.Error()
	}
}
