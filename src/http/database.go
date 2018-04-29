package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"log"
)

var connection *sql.DB

// Widget :
type Widget struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

/** */
func initDB(dbname, username, password string) {

	conn, err := sql.Open("mysql", username+":"+password+"@/"+dbname)
	if err != nil {
		log.Fatal(err)
	}

	connection = conn
}

/** */
func closeDB() {
	connection.Close()
}

/** */
func checkAuthDB(name, pass string) bool {

	stmt, err := connection.Prepare("select name from user where name=? and pass=?")

	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRow(name, hashPassword(pass))

	err = row.Scan(&name, &pass)

	return err != sql.ErrNoRows
}

/** */
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
