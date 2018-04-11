package main

import (
	"rpc_test/shared"
)

//Reverse :
func (client *Client) Reverse(data string) string {

	encryptedData, err := shared.Encrypt(shared.TestKey, []byte(data))

	if err != nil {
		panic(err)
	}

	in := &shared.Request{Body: encryptedData}

	var out shared.Response

	err = client.conn.Call("StringOps.Reverse", in, &out)

	if err != nil {
		panic(err)
	}

	decryptedData, err := shared.Decrypt(shared.TestKey, out.Body)

	if err != nil {
		panic(err)
	}
	return string(decryptedData)
}

// ToUpper :
func (client *Client) ToUpper(data string) string {

	encryptedData, err := shared.Encrypt(shared.TestKey, []byte(data))

	if err != nil {
		panic(err)
	}

	in := &shared.Request{Body: encryptedData}

	var out shared.Response

	err = client.conn.Call("StringOps.ToUpper", in, &out)

	if err != nil {
		panic(err)
	}

	decryptedData, err := shared.Decrypt(shared.TestKey, out.Body)

	if err != nil {
		panic(err)
	}
	return string(decryptedData)

}

// ToLower :
func (client *Client) ToLower(data string) string {

	encryptedData, err := shared.Encrypt(shared.TestKey, []byte(data))

	if err != nil {
		panic(err)
	}

	in := &shared.Request{encryptedData}

	var out shared.Response

	err = client.conn.Call("StringOps.ToLower", in, &out)

	if err != nil {
		panic(err)
	}

	decryptedData, err := shared.Decrypt(shared.TestKey, out.Body)

	if err != nil {
		panic(err)
	}
	return string(decryptedData)

}
