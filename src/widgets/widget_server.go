package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
	database Database
}

/**  */
func setJsonContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

/**  */
func (server *Server) getWidgets(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	log.Printf("getWidgets:")

	json.NewEncoder(w).Encode(server.database.getDBWidgets())
}

/**  */
func (server *Server) getWidget(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err == nil {
		log.Printf("getWidget: %d\n", id)

		json.NewEncoder(w).Encode(server.database.getDBWidget(id))
	} else {
		log.Fatal(err)
	}
}

/**  */
func (server *Server) deleteWidget(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err == nil {
		log.Printf("deleteWidget: %d\n", id)
	} else {
		log.Fatal(err)
	}

	server.database.deleteDBWidget(id)
}

/**  */
func (server *Server) addWidget(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	log.Printf("addWidget:\n")

	widget := Widget{}

	_ = json.NewDecoder(r.Body).Decode(&widget)

	server.database.addDBWidget(widget)
}

/**  */
func (server *Server) updateWidget(w http.ResponseWriter, r *http.Request) {
	setJsonContentType(w)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err == nil {
		log.Printf("updateWidget: %d\n", id)
	} else {
		log.Fatal(err)
	}
	//	json.NewEncoder(w).Encode(widgets)
}

/**  */
func main() {
	fmt.Println("main")

	server := Server{}

	server.database.Initialise("", "", "")

	r := mux.NewRouter()

	r.HandleFunc("/widgets/", server.getWidgets).Methods("GET")
	r.HandleFunc("/widgets/", server.addWidget).Methods("POST")
	r.HandleFunc("/widget/{id}", server.getWidget).Methods("GET")
	r.HandleFunc("/widget/{id}", server.updateWidget).Methods("PUT")
	r.HandleFunc("/widget/{id}", server.deleteWidget).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
