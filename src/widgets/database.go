package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	connection *sql.DB
}

/** */
func (database *Database) Initialise(user, password, dbname string) {

	conn, err := sql.Open("mysql", "root:root@/test")

	if err != nil {
		log.Fatal(err)
	}

	database.connection = conn
}

/** */
func (database *Database) getDBWidgets() []Widget {

	stmt, err := database.connection.Prepare("select id,name,color from widget")

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
		if err := rows.Scan(&widget.Id, &widget.Name, &widget.Color); err != nil {
			log.Fatal(err)
		}
		widgets = append(widgets, widget)
	}
	return widgets
}

/** */
func (database *Database) getDBWidget(id int) Widget {

	stmt, err := database.connection.Prepare("select id,name,color from widget where id =?")

	if err != nil {
		log.Fatal(err)
	}

	var widget Widget

	err = stmt.QueryRow(id).Scan(&widget.Id, &widget.Name, &widget.Color)

	if err != nil {
		log.Fatal(err)
	}

	return widget
}

/** */
func (database *Database) deleteDBWidget(id int) {

	stmt, err := database.connection.Prepare("delete from widget where id=?")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}
}

/** */
func (database *Database) addDBWidget(widget Widget) {

	stmt, err := database.connection.Prepare("insert into widget values (?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(widget.Id, widget.Name, widget.Color)

	if err != nil {
		log.Fatal(err)
	}
}
