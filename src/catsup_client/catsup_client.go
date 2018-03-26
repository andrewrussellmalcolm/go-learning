package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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

// BASE_URL :
const BASE_URL = "https://localhost:8443/catsup/"

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

	var users []User
	var messages []Message

	for true {
		fmt.Printf("===================================================\n")
		fmt.Printf("Enter an option\n")
		fmt.Printf("u = list users\n")
		fmt.Printf("m = list messages\n")
		fmt.Printf("s = send message\n")
		fmt.Printf("q = quit\n")
		fmt.Printf("===================================================\n")
		fmt.Print(">")
		line, _ := reader.ReadString('\n')

		var cmd rune
		_, err := fmt.Sscanf(line, "%c", &cmd)

		if err == nil {
			switch cmd {

			case 'q':
				os.Exit(0)

			case 'u':
				users = listUsers(name, pass)
				for n, user := range users {
					fmt.Println(n+1, user.Name, user.Email)
				}

			case 'm':
				messages = listMessages(name, pass)

				for n, message := range messages {
					fmt.Println(n+1, message.Timestamp.Format("3:04PM"), ": ", message.Text)
				}

			case 's':

				var userIndex string
				_, err = fmt.Sscanf(line, "%c %s", &cmd, &userIndex)

				if err != nil {
					fmt.Println("Not enough data")
					panic(err)
				}

				offset := strings.LastIndex(line, userIndex) + len(userIndex)
				fmt.Println(cmd, userIndex, line[offset:])

				index, err := strconv.Atoi(userIndex)

				if err != nil {
					fmt.Println("user number must be an integer")
					panic(err)
				}

				userID := users[index-1].ID.Hex()

				sendMessage(userID, name, pass, line[offset:])
			}
		} else {
			panic(err)
		}
	}
}

// sendMessage
func sendMessage(userID, name, pass, text string) {
	fmt.Printf("sending %s to %s\n", text, userID)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// send to andrew
	url := BASE_URL + "sendmessage/" + userID

	message := Message{}
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
func listMessages(name, pass string) []Message {

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	url := BASE_URL + "getmessagelist/5ab8a3387ebb035f8933ef6d"

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
	var messages []Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		panic(err)
	}

	return messages
}

// listUsers
func listUsers(name, pass string) []User {

	//	httpClient := http.DefaultClient

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url := BASE_URL + "getuserlist"

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

	return users
}
