//
// see: https://gist.github.com/denji/12b3a568f092ab951456 for ssl key generation
//
package main

import (
	"catsup/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/globalsign/mgo/bson"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

/**  */
func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

/**  */
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])

	//	err = deleteMessageDB(id)

	if err == nil {
		json.NewEncoder(w).Encode(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

/**  */
func sendMessage(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)

	session, _ := store.Get(r, "cookie-name")

	user := session.Values["user"].(*database.User)
	var message database.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	toID := bson.ObjectIdHex(mux.Vars(r)["to_id"])

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
	json.NewEncoder(w).Encode(users)
}

func getMessageList(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "cookie-name")

	user := session.Values["user"].(*database.User)

	fromID := bson.ObjectIdHex(mux.Vars(r)["from_id"])

	messages := database.GetMessageList(user.ID, fromID)
	json.NewEncoder(w).Encode(messages)
}

//
func temp(w http.ResponseWriter, r *http.Request) {

	log.Printf("user logged in\n")
	w.WriteHeader(http.StatusOK)

	session, _ := store.Get(r, "cookie-name")

	user := session.Values["user"].(*database.User)

	fmt.Println(user)
}

// checkAuth
func checkAuth(name, pass string, r *http.Request) bool {

	session, _ := store.Get(r, "cookie-name")

	user := database.CheckAuth(name, pass)

	if user != nil {
		session.Values["authenticated"] = true
		session.Values["user"] = user
		return true
	}

	session.Values["authenticated"] = false
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

	router.HandleFunc("/catsup/temp", temp).Methods("GET")
	router.HandleFunc("/catsup/getuserlist", getUserList).Methods("GET")
	router.HandleFunc("/catsup/getmessagelist/{from_id}", getMessageList).Methods("GET")
	router.HandleFunc("/catsup/sendmessage/{to_id}", sendMessage).Methods("POST")
	//router.HandleFunc("/Messcatsupage/{id}", updateMessage).Methods("PUT")
	router.HandleFunc("/catsup/{id}", deleteMessage).Methods("DELETE")

	router.Use(httpauth.BasicAuth(httpauth.AuthOptions{AuthFunc: checkAuth}))
	log.Fatal(http.ListenAndServeTLS(config.Server.Port, config.Server.CertPath, config.Server.KeyPath, router))
}
