package main

import (
	"bytes"
	"catsup/shared"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

var serverURL string

var httpClient *http.Client

func initClient(url string) {

	serverURL = url
	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// sendMessage
func sendMessage(userID, name, pass, text string) {
	fmt.Printf("sending %s to %s\n", text, userID)

	// send to andrew
	url := serverURL + "message?to_id=" + userID

	message := shared.Message{}
	message.Text = text

	val, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(val))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(name, pass)

	req.Header.Set("Accept", "application/json")

	_, err = httpClient.Do(req)
	if err != nil {
		panic(err)
	}
}

// listMessages
func listMessages(userID, name, pass string) []shared.Message {

	url := serverURL + "messagelist?from_id=" + userID

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(name, pass)

	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var messages []shared.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		panic(err)
	}

	return messages
}

// listNewMessages
func listNewMessages(userID, name, pass string) []shared.Message {

	url := serverURL + "newmessagelist?from_id=" + userID

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(name, pass)

	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var messages []shared.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		panic(err)
	}

	return messages
}

// listUsers
func listUsers(name, pass string) []shared.User {

	url := serverURL + "userlist"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(name, pass)

	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var users []shared.User
	err = json.NewDecoder(resp.Body).Decode(&users)

	if err != nil {
		panic(err)
	}

	return users
}

// getUserStatus
func getUserStatus(userID, name, pass string) shared.UserStatus {

	url := serverURL + "userstatus?user_id=" + userID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(name, pass)

	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var userStatus shared.UserStatus
	err = json.NewDecoder(resp.Body).Decode(&userStatus)

	if err != nil {
		panic(err)
	}

	return userStatus
}

// getUserStatus
func updateUserStatus(name, pass string) {

	url := serverURL + "userstatus"

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(name, pass)

	req.Header.Set("Accept", "application/json")

	_, err = httpClient.Do(req)
	if err != nil {
		panic(err)
	}
}
