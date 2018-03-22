//
// see: https://gist.github.com/denji/12b3a568f092ab951456 for ssl key generation
//
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

/**  */
func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

/**  */
func getWidgets(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	widgets := getDBWidgets()
	json.NewEncoder(w).Encode(widgets)
}

/**  */
func getWidget(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	widget, err := getDBWidget(id)

	if err == nil {
		json.NewEncoder(w).Encode(widget)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

/**  */
func deleteWidget(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	err = deleteDBWidget(id)

	if err == nil {
		json.NewEncoder(w).Encode(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

/**  */
func addWidget(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)

	var widget Widget

	err := json.NewDecoder(r.Body).Decode(&widget)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	err = addDBWidget(widget)

	if err == nil {
		json.NewEncoder(w).Encode(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

/**  */
func updateWidget(w http.ResponseWriter, r *http.Request) {
	setJSONContentType(w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	var widget Widget
	err = json.NewDecoder(r.Body).Decode(&widget)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	widget.ID = id
	err = updateDBWidget(widget)

	if err == nil {
		json.NewEncoder(w).Encode(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
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

	log.SetOutput(f)
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

	// start the router
	router := mux.NewRouter()

	router.HandleFunc("/widgets", getWidgets).Methods("GET")
	router.HandleFunc("/widgets", addWidget).Methods("POST")
	router.HandleFunc("/widget/{id}", getWidget).Methods("GET")
	router.HandleFunc("/widget/{id}", updateWidget).Methods("PUT")
	router.HandleFunc("/widget/{id}", deleteWidget).Methods("DELETE")

	router.Use(httpauth.BasicAuth(httpauth.AuthOptions{AuthFunc: checkAuth}))
	log.Fatal(http.ListenAndServeTLS(config.Server.Port, config.Server.CertPath, config.Server.KeyPath, router))
}
