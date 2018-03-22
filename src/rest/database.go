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
func addDBWidget(widget Widget) error {

	stmt, err := connection.Prepare("insert into widget (name,color) values (?,?)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(widget.Name, widget.Color)

	return err
}

/** */
func updateDBWidget(widget Widget) error {

	stmt, err := connection.Prepare("update widget set name=?, color=?  where id=?")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(widget.Name, widget.Color, widget.ID)

	return err
}

/** */
func getDBWidgets() []Widget {

	stmt, err := connection.Prepare("select id,name,color from widget")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	widgets := []Widget{}

	for rows.Next() {
		var widget Widget
		if err := rows.Scan(&widget.ID, &widget.Name, &widget.Color); err != nil {
			log.Fatal(err)
		}
		widgets = append(widgets, widget)
	}
	return widgets
}

/** */
func getDBWidget(id int) (Widget, error) {

	stmt, err := connection.Prepare("select id,name,color from widget where id=?")

	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRow(id)

	widget := Widget{}

	err = row.Scan(&widget.ID, &widget.Name, &widget.Color)

	if err == sql.ErrNoRows {
		return widget, err
	}

	return widget, nil
}

/** */
func deleteDBWidget(id int) error {

	stmt, err := connection.Prepare("delete from widget where id=?")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
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
