package main

import "fmt"
import "github.com/ktokarz289/Cryptopal/Chapters/chapter1"

func main() {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	value := chapter1.ConvertHexToBase64(hex)
	fmt.Println(value)
}
