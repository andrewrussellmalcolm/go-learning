package main

import (
	"bufio"
	"bytes"
	"catsup/shared"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/globalsign/mgo/bson"
	//	termbox "github.com/nsf/termbox-go"
)

// baseURL :
//const baseURL = "https://35.176.163.49:443/catsup/"
const baseURL = "https://localhost:8443/catsup/"

var httpClient *http.Client

// main :
func main() {
	if len(os.Args) != 3 {

		fmt.Printf("Usage %s useranme password\n", os.Args[0])
		os.Exit(0)
	}

	// err := termbox.Init()
	// if err != nil {
	// 	panic(err)
	// }
	// defer termbox.Close()

	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	name := os.Args[1]
	pass := os.Args[2]

	reader := bufio.NewReader(os.Stdin)

	var users []shared.User

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		for {
			updateUserStatus(name, pass)
			time.Sleep(30 * time.Second)
		}
	}()

	for true {
		///termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		fmt.Printf("===================================================\n")
		fmt.Printf("Enter an option\n")
		fmt.Printf("lu = list users\n")
		fmt.Printf("us = user status\n")
		fmt.Printf("lm = list messages (e.g m 1\n")
		fmt.Printf("sm = send message (e.g. s 1 this is my message)\n")
		fmt.Printf("q = quit\n")
		fmt.Printf("===================================================\n")
		fmt.Print(">")
		line, _ := reader.ReadString('\n')

		words := strings.Split(strings.TrimSpace(line), " ")

		if len(words) > 0 {
			switch words[0] {

			case "q":
				os.Exit(0)

			case "lu":
				users = listUsers(name, pass)
				printUsers(users)

			case "lm":
				if len(words) > 1 {

					index, err := strconv.Atoi(words[1])
					index--

					if err != nil {
						fmt.Println("user number must be an integer")
						break
					}

					if index < 0 || index >= len(users) {
						fmt.Println("user number out of range")
						break

					}
					userID := users[index].ID.Hex()

					printMessages(users[index].ID, listMessages(userID, name, pass))
				} else {
					fmt.Println("no user id supplied")
				}

			case "sm":

				if len(words) > 1 {

					index, err := strconv.Atoi(words[1])
					index--

					if err != nil {
						fmt.Println("user number must be an integer")
						break
					}

					if index < 0 || index >= len(users) {
						fmt.Println("user number out of range")
						break
					}

					userID := users[index].ID.Hex()

					sendMessage(userID, name, pass, strings.Join(words[2:], " "))
				} else {
					fmt.Println("no user id or message supplied")
				}
			case "us":
				if len(words) > 1 {

					index, err := strconv.Atoi(words[1])
					index--

					if err != nil {
						fmt.Println("user number must be an integer")
						break
					}

					if index < 0 || index >= len(users) {
						fmt.Println("user number out of range")
						break

					}
					userID := users[index].ID.Hex()

					printUserStatus(getUserStatus(userID, name, pass))

				} else {
					fmt.Println("no user id supplied")
				}

			default:
				fmt.Println("unknown command")
			}
		}
	}
}

// sendMessage
func sendMessage(userID, name, pass, text string) {
	fmt.Printf("sending %s to %s\n", text, userID)

	// send to andrew
	url := baseURL + "message?to_id=" + userID

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

	url := baseURL + "messagelist?from_id=" + userID

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

	url := baseURL + "userlist"

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

	url := baseURL + "userstatus?user_id=" + userID

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

	url := baseURL + "userstatus"

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

func printMessages(user bson.ObjectId, messages []shared.Message) {
	fmt.Println("=========== FROM ============== TO =========")
	for n, message := range messages {

		if user == message.From {
			fmt.Println(n+1, message.Timestamp.Format("3:04PM"), ": ", "\t\t\t"+message.Text)
		} else {

			fmt.Println(n+1, message.Timestamp.Format("3:04PM"), ": ", message.Text)
		}
	}
}

func printUsers(users []shared.User) {
	for n, user := range users {
		fmt.Println(n+1, user.Name, user.Email, user.Timestamp)
	}
}

func printUserStatus(userStatus shared.UserStatus) {

	if userStatus == shared.ONLINE {
		fmt.Println("User online")
	} else {
		fmt.Println("User offline")
	}

}
