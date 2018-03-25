package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

//
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func main() {

	fmt.Println(hashPassword(os.Args[1]))
}
