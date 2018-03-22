package main

import (
	"fmt"
	"rpc_test/shared"
	"strings"
)

// StringOps :
type StringOps struct {
	callCount int64
}

// Reverse : reverse string
func (stringOps *StringOps) Reverse(in *shared.Request, out *shared.Response) error {

	decryptedData, err := shared.Decrypt(shared.TestKey, in.Body)

	if err != nil {
		return err
	}

	r := []rune(string(decryptedData))
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	encryptedData, err := shared.Encrypt(shared.TestKey, []byte(string(r)))

	if err != nil {
		return err
	}

	out.Body = encryptedData

	stringOps.callCount++
	fmt.Printf("Call count %d\r", stringOps.callCount)

	return nil
}

// ToUpper :
func (stringOps *StringOps) ToUpper(in *shared.Request, out *shared.Response) error {

	decryptedData, err := shared.Decrypt(shared.TestKey, in.Body)

	if err != nil {
		return err
	}

	u := strings.ToUpper(string(decryptedData))

	encryptedData, err := shared.Encrypt(shared.TestKey, []byte(string(u)))

	if err != nil {
		return err
	}

	out.Body = encryptedData

	stringOps.callCount++
	fmt.Printf("Call count %d\r", stringOps.callCount)
	return nil
}

// ToLower :
func (stringOps *StringOps) ToLower(in *shared.Request, out *shared.Response) error {

	decryptedData, err := shared.Decrypt(shared.TestKey, in.Body)

	if err != nil {
		return err
	}

	l := strings.ToLower(string(decryptedData))

	encryptedData, err := shared.Encrypt(shared.TestKey, []byte(string(l)))

	if err != nil {
		return err
	}

	out.Body = encryptedData

	stringOps.callCount++
	fmt.Printf("Call count %d\r", stringOps.callCount)
	return nil
}
