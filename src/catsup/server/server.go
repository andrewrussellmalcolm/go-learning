//
// see: https://gist.github.com/denji/12b3a568f092ab951456 for ssl key generation
//
package main

import (
	"catsup/server/database"
	"catsup/shared"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

/**  */
func main() {

	// read the configuration
	config := readConfig()

	// start logging
	logfile := initLogging()
	defer logfile.Close()

	// start the database
	err := database.Init(config.Database.DBName)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// inittialise the cookie store
	store = sessions.NewCookieStore([]byte(config.Server.CookieKey))

	// start the router
	router := mux.NewRouter()

	router.HandleFunc("/catsup/userstatus", getUserStatus).Methods("GET")
	router.HandleFunc("/catsup/userlist", getUserList).Methods("GET")
	router.HandleFunc("/catsup/messagelist", getMessageList).Methods("GET")
	router.HandleFunc("/catsup/message", postMessage).Methods("POST")
	router.HandleFunc("/catsup/message", updateMessage).Methods("PUT")
	router.HandleFunc("/catsup/message", deleteMessage).Methods("DELETE")

	router.Use(httpauth.BasicAuth(httpauth.AuthOptions{AuthFunc: checkAuth}))
	log.Fatal(http.ListenAndServeTLS(config.Server.Port, config.Server.CertPath, config.Server.KeyPath, router))
}

/**  */
func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

/**  */
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
}

/**  */
func updateMessage(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
}

/**  */
func postMessage(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	queryValues := r.URL.Query()

	session, _ := store.Get(r, "session")

	user := session.Values["user"].(*shared.User)
	var message shared.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	toID := bson.ObjectIdHex(queryValues.Get("to_id"))

	_, err = database.InsertMessage(message, toID, user.ID)

	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
	}
}

// getUserList :
func getUserList(w http.ResponseWriter, r *http.Request) {

	users := database.GetUserList()

	if users != nil {
		json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// getUserStatus :
func getUserStatus(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	fromID := queryValues.Get("from_id")
	user := database.GetUserByID(fromID)
	lastAccess := user.LastAccess
	json.NewEncoder(w).Encode(lastAccess)
}

func getMessageList(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)

	queryValues := r.URL.Query()
	session, _ := store.Get(r, "session")
	toID := session.Values["user"].(*shared.User).ID
	fromID := bson.ObjectIdHex(queryValues.Get("from_id"))
	messages := database.GetMessageList(toID, fromID)

	if messages != nil {
		json.NewEncoder(w).Encode(messages)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// checkAuth
func checkAuth(name, pass string, r *http.Request) bool {

	session, _ := store.Get(r, "session")

	user := database.CheckAuth(name, pass)

	if user != nil {
		session.Values["user"] = user

		user.LastAccess = time.Now()
		database.UpdateUserAccessTime(user.ID, time.Now())
		return true
	}

	log.Printf("Attempted access by %s failed\n", name)
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
