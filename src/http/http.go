//
// see: https://gist.github.com/denji/12b3a568f092ab951456 for ssl key generation
//
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

var currentFrame []byte

/**  */
func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write([]byte("hello world"))

	fmt.Println(len(currentFrame))

	w.Write(currentFrame)

	w.WriteHeader(http.StatusOK)
}

/** */
func checkAuth(user, pass string, r *http.Request) bool {
	if checkAuthDB(user, pass) {
		return true
	}

	log.Printf("Attempted access by %s failed\n", user)
	return false
}

/** */
func initLogging() *os.File {

	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v\n", err))
	}

	//log.SetOutput(f)
	return f
}

/**  */
func main() {

	// read the configuration
	config := readConfig()

	// start logging
	logfile := initLogging()
	defer logfile.Close()

	// start the database
	initDB(config.Database.DBName, config.Database.Username, config.Database.Password)
	defer closeDB()

	initWebcam()
	defer closeWebcam()

	listWebcamFormats()

	frame := make(chan []byte)

	go func() {

		for {
			currentFrame = <-frame

		}
	}()

	go startCapture(frame)

	// start the router
	router := mux.NewRouter()

	router.HandleFunc("/hello", handleHello).Methods("GET")

	router.Use(httpauth.BasicAuth(httpauth.AuthOptions{AuthFunc: checkAuth}))
	log.Fatal(http.ListenAndServeTLS(config.Server.Port, config.Server.CertPath, config.Server.KeyPath, router))
}
