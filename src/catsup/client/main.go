package main

import (
	"bufio"
	"catsup/shared"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/globalsign/mgo/bson"
)

// main :
func main() {

	if len(os.Args) != 3 {

		fmt.Printf("Usage %s useranme password\n", os.Args[0])
		os.Exit(0)
	}

	config := readConfig()
	fmt.Println("Connecting to: " + config.Server.ServerURL)

	initClient(config.Server.ServerURL)

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
		fmt.Printf("===================================================\n")
		fmt.Printf("Enter an option\n")
		fmt.Printf("lu = list users\n")
		fmt.Printf("us = user status\n")
		fmt.Printf("lm = list messages (e.g lm 1)\n")
		fmt.Printf("sm = send message (e.g. sm 1 this is my message)\n")
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

func printMessages(user bson.ObjectId, messages []shared.Message) {
	fmt.Println("============ TO =============== FROM ==============")
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
