package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo/bson"
)

// Status :
type Status int

const (
	// WAITING :
	WAITING = 0
	// SENT :
	SENT = 1
	//RECEIVED :
	RECEIVED = 2
	// READ :
	READ = 3
)

// Message :
type Message struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Text      string
	Timestamp time.Time
	To        bson.ObjectId
	From      bson.ObjectId
	Status    Status
}

// User :
type User struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Email string
	Hash  string
}

func main() {
	if len(os.Args) != 3 {

		fmt.Printf("Usage %s useranme password\n", os.Args[0])
		os.Exit(0)
	}

	//stop := make(chan os.Signal, 1)
	//signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	name := os.Args[1]
	pass := os.Args[2]

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Printf("===================================================\n")
		fmt.Printf("Enter an option\n")
		fmt.Printf("u = list users\n")
		fmt.Printf("l = list messages\n")
		fmt.Printf("s = send message\n")
		fmt.Printf("q = quit\n")
		fmt.Printf("===================================================\n")
		fmt.Print(">")
		text, _ := reader.ReadString('\n')

		var cmd rune
		_, err := fmt.Sscanf(text, "%c", &cmd)

		if err == nil {
			switch cmd {

			case 'q':
				os.Exit(0)

			case 'u':
				listUsers(name, pass)

			case 'l':
				listMessages(name, pass)

			case 's':

				var user, text string
				_, err = fmt.Sscanf(text, "%c %s %20[^\t\n]", &cmd, &user, &text)

				if err == nil {
					sendMessage(user, name, pass, text)
				} else {
					fmt.Println("Not enough data")
				}

			}
		} else {
			panic(err)
		}
	}

}

// sendMessage
func sendMessage(user, name, pass, text string) {
	fmt.Println(user, text)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// send to andrew
	url := "https://localhost:8080/catsup/sendmessage/5ab5ff1e89aba31d84acc2c4"

	message := Message{}
	message.Text = text

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
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
	var messages []Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		panic(err)
	}

}

// listMessages
func listMessages(name, pass string) {

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	url := "https://localhost:8080/catsup/getmessagelist/5ab5ff1e89aba31d84acc2c4"

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
	var messages []Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		fmt.Println(message.Timestamp.Format("3:04PM"), ": ", message.Text)
	}

}

// listUsers
func listUsers(name, pass string) {

	//	httpClient := http.DefaultClient

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	url := "https://localhost:8080/catsup/getuserlist"

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
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)

	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.Name, user.Email)
	}
}
