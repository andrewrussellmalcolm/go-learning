package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

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

/**  */
func main() {

	initDB("", "", "")
	defer closeDB()

	r := mux.NewRouter()

	r.HandleFunc("/widgets", getWidgets).Methods("GET")
	r.HandleFunc("/widgets", addWidget).Methods("POST")
	r.HandleFunc("/widget/{id}", getWidget).Methods("GET")
	r.HandleFunc("/widget/{id}", updateWidget).Methods("PUT")
	r.HandleFunc("/widget/{id}", deleteWidget).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
