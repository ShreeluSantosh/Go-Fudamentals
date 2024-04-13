package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func getSHA1Hash(message string) string {
	hash := sha1.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

func getSHA256Hash(message string) string {
	hash := sha256.Sum224([]byte(message))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println("Input string: ")
	var input string
	fmt.Scan(&input)
	fmt.Println("SHA1 Hashed value: ", getSHA1Hash(input))
	fmt.Println("SHA256 Hashed value: ", getSHA256Hash(input))

}
