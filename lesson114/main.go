package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {

	str := "Hello,World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encoded)

	data, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(data))

	std := "Hello world!"
	hash := sha256.Sum256([]byte(std))
	hashStd := fmt.Sprintf("%x", hash)

	fmt.Println(hashStd)
}
