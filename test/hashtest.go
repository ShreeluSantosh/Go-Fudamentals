package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println("Input string: ")
	var input string
	fmt.Scan(&input)
	fmt.Println("MD5 Hashed value: ", getMD5Hash(input))

}
