package main

import (
	"catsup/shared"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var users []shared.User

// main :
func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Usage %s useranme password\n", os.Args[0])
		os.Exit(0)
	}

	termbox.Init()
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	config := readConfig()

	initClient(config.Server.ServerURL)

	name := os.Args[1]
	pass := os.Args[2]

	go func() {

		for {
			updateUserStatus(name, pass)
			time.Sleep(30 * time.Second)
		}
	}()

	tbclear()

exit:
	for true {

		printMenu(config.Server.ServerURL)

		termbox.Flush()

		var line strings.Builder
		termbox.SetCursor(1, 10)

	next:
		for {

			switch ev := termbox.PollEvent(); ev.Type {

			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					break exit
				case termbox.KeyEnter:
					termbox.HideCursor()
					processLine(name, pass, line.String())
					break next
				case termbox.KeySpace:
					line.WriteRune(' ')
					tbprint(1, 10, termbox.ColorGreen, termbox.ColorBlack, line.String())
					termbox.Flush()
					termbox.SetCursor(line.Len(), 10)
				default:
					if ev.Ch != 0 {
						line.WriteRune(ev.Ch)
						tbprint(1, 10, termbox.ColorGreen, termbox.ColorBlack, line.String())
						termbox.Flush()
						termbox.SetCursor(line.Len(), 10)
					}
				}
			}
		}
	}
}

func processLine(name, pass, line string) {

	tbclear()

	words := strings.Split(strings.TrimSpace(line), " ")

	if len(words) > 0 {
		switch words[0] {

		case "q":
			os.Exit(0)

		case "lu":
			users = listUsers(name, pass)
			printUsers(0, 12, users)

		case "lm":
			if len(words) > 1 {

				index, err := strconv.Atoi(words[1])
				index--

				if err != nil {
					tbprint(0, 0, termbox.ColorRed, termbox.ColorWhite, "user number must be an integer")
					break
				}

				if index < 0 || index >= len(users) {
					tbprint(0, 0, termbox.ColorRed, termbox.ColorWhite, "user number out of range")
					break

				}
				userID := users[index].ID.Hex()

				printMessages(0, 12, users[index].ID, listMessages(userID, name, pass))
			} else {
				tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "no user id supplied")
			}

		case "nm":
			if len(words) > 1 {

				index, err := strconv.Atoi(words[1])
				index--

				if err != nil {
					tbprint(0, 0, termbox.ColorRed, termbox.ColorWhite, "user number must be an integer")
					break
				}

				if index < 0 || index >= len(users) {
					tbprint(0, 0, termbox.ColorRed, termbox.ColorWhite, "user number out of range")
					break

				}
				userID := users[index].ID.Hex()

				printMessages(0, 12, users[index].ID, listNewMessages(userID, name, pass))
			} else {
				tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "no user id supplied")
			}

		case "sm":

			if len(words) > 1 {

				index, err := strconv.Atoi(words[1])
				index--

				if err != nil {
					tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "user number must be an integer")
					break
				}

				if index < 0 || index >= len(users) {
					tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "user number out of range")
					break
				}

				userID := users[index].ID.Hex()

				sendMessage(userID, name, pass, strings.Join(words[2:], " "))
			} else {
				tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "no user id or message supplied")
			}
		case "us":
			if len(words) > 1 {

				index, err := strconv.Atoi(words[1])
				index--

				if err != nil {
					tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "user number must be an integer")
					break
				}

				if index < 0 || index >= len(users) {
					tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "user number out of range")
					break

				}
				userID := users[index].ID.Hex()

				printUserStatus(getUserStatus(userID, name, pass))

			} else {
				tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "no user id supplied")
			}

		default:
			tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "unknown command")
		}
	} else {
		tbprint(0, 12, termbox.ColorRed, termbox.ColorWhite, "unknown error")
	}
}

func printMessages(x, y int, user bson.ObjectId, messages []shared.Message) {

	tbprint(x+0, y, termbox.ColorWhite, termbox.ColorBlack, "USER")
	tbprint(x+6, y, termbox.ColorWhite, termbox.ColorBlack, "TIME")
	tbprint(x+25, y, termbox.ColorWhite, termbox.ColorBlack, "STATUS")
	tbprint(x+35, y, termbox.ColorWhite, termbox.ColorBlack, "FROM")
	tbprint(x+60, y, termbox.ColorWhite, termbox.ColorBlack, "TO")
	y++

	for n, message := range messages {
		tbprint(x, y, termbox.ColorGreen, termbox.ColorBlack, fmt.Sprintf("%d", n+1))
		tbprint(x+6, y, termbox.ColorGreen, termbox.ColorBlack, fmt.Sprintf("%s", message.Timestamp.Format("Mon Jan 2 3:04PM")))
		tbprint(x+25, y, termbox.ColorGreen, termbox.ColorBlack, fmt.Sprintf("%d", message.Status))

		if user == message.From {
			tbprint(x+60, y, termbox.ColorCyan, termbox.ColorBlack, message.Text)
		} else {
			tbprint(x+35, y, termbox.ColorCyan, termbox.ColorBlack, message.Text)
		}

		y++
	}
	termbox.Flush()
}

func printUsers(x, y int, users []shared.User) {
	y++
	for n, user := range users {
		tbprint(x, y, termbox.ColorGreen, termbox.ColorBlack, fmt.Sprintf("%d %s%s %v", n+1, user.Name, user.Email, user.Timestamp))
		y++
	}
	termbox.Flush()
}

func printMenu(serverURL string) {
	tbprint(0, 1, termbox.ColorYellow, termbox.ColorBlack, "Connected to: "+serverURL)
	tbprint(0, 3, termbox.ColorGreen, termbox.ColorBlack, "Enter an option")
	tbprint(0, 4, termbox.ColorGreen, termbox.ColorBlack, "lu = list users")
	tbprint(0, 5, termbox.ColorGreen, termbox.ColorBlack, "us = user status")
	tbprint(0, 6, termbox.ColorGreen, termbox.ColorBlack, "lm = list messages (e.g lm 1)")
	tbprint(0, 6, termbox.ColorGreen, termbox.ColorBlack, "nm = list new messages (e.g nm 1)")
	tbprint(0, 7, termbox.ColorGreen, termbox.ColorBlack, "sm = send message (e.g. sm 1 this is my message)")
	tbprint(0, 8, termbox.ColorGreen, termbox.ColorBlack, "q = quit")
	tbprint(0, 10, termbox.ColorGreen, termbox.ColorBlack, ">                                              ")
}

func printUserStatus(userStatus shared.UserStatus) {
	if userStatus == shared.ONLINE {
		tbprint(0, 12, termbox.ColorGreen, termbox.ColorBlack, "User online")
	} else {
		tbprint(0, 12, termbox.ColorGreen, termbox.ColorBlack, "User offline")
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func tbclear() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
}
