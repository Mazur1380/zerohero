package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	x := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UifQ.7TWR5jPkezkACsSMVnUEvuSRzvfFcycBFkh-EOwoCzs"
	y := strings.Split(x, ".")

	decodedBytes, err := base64.StdEncoding.DecodeString(y[0])
	if err != nil {
		fmt.Println("Ошибка при декодировании:", err)
		return
	}
	fmt.Println(string(decodedBytes))
	c := y[1]
	for len(c)%4 != 0 {
		c += "="
	}

	decodedBytes, err = base64.StdEncoding.DecodeString(c)
	if err != nil {
		fmt.Println("Ошибка при декодировании:", err)
		return
	}

	fmt.Println(string(decodedBytes))

	secretKey := "1380"
	message := y[0] + "." + y[1]

	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(message))
	sum := hash.Sum(nil)
	sig := base64.StdEncoding.EncodeToString(sum)
	fmt.Println(sig)

	// fmt.Println(hex.EncodeToString(sum))
	// fmt.Println(string(sum))
}
