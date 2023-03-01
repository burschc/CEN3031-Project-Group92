package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func SetupOrOpenBasicDatabase() {
	//Open the database.
	database, _ := sql.Open("sqlite3", "testdb.db")
	//Make the database if it doesn't exist.
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, passtype INTEGER)")
	statement.Exec()
	//Prepare the to add a user to the database, then actually do it.
	statement, _ = database.Prepare("INSERT INTO users (username, password, passtype) VALUES (?, ?, ?)")
	statement.Exec("TestUser", "TestPassword", 1)
	//Prepare to print the entire list of entries.
	rows, _ := database.Query("SELECT id, username, password, passtype FROM users")
	var id int
	var username string
	var password string
	var passtype int
	for rows.Next() {
		rows.Scan(&id, &username, &password, &passtype)
		fmt.Println(strconv.Itoa(id) + ": " + username + " " + password + " " + strconv.Itoa(passtype))
	}
}
