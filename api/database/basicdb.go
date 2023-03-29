package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"html"
	"log"
	"net/http"
	_ "os"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	//"google.golang.org/appengine"
)

// DatabaseName is the filename of the database which stores user account information.
const DatabaseName = "accountDB"

var database *sql.DB

func DatabaseHandlers(r *mux.Router) {
	r.HandleFunc("/signup", signup).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/login", login).Methods(http.MethodPost, http.MethodOptions)

	log.Print("Registered database handlers.")
}

func DeclareDatabase(dbname string) {
	//Open the database.
	database, _ = sql.Open("sqlite3", dbname)
	//Make the database if it doesn't exist.
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, passtype INTEGER)")
	_, _ = statement.Exec()
}

func LegacySetupOrOpenBasicDatabase() {
	//Open the database.
	database, _ = sql.Open("sqlite3", "testdb.db")
	//Make the database if it doesn't exist.
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, passtype INTEGER)")
	_, _ = statement.Exec()
	//Prepare the to add a user to the database, then actually do it.
	statement, _ = database.Prepare("INSERT INTO users (username, password, passtype) VALUES (?, ?, ?)")
	_, _ = statement.Exec("TestUser", "TestPassword", 1)
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

// printDatabase is a test function which prints the contents of the database.
func printDatabase() {
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

// signup adds a specified account with given username and password to the database if it is not already present.
func signup(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	username := html.EscapeString(req.FormValue("username"))
	password := html.EscapeString(req.FormValue("password"))

	var user string

	err := database.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	log.Println(err)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Hash error, unable to create your account.", 500)
			return
		}
		//passtype -1 being "no passtype declared" -> defaulting to showing all types
		if _, err = database.Exec("INSERT INTO users (username, password, passtype) VALUES(?, ?, ?)", username, hashedPassword, -1); err != nil {
			http.Error(res, "Insert error, unable to create your account.", 500)
			return
		}

		res.Write([]byte("User created!"))
		return
	case err != nil:
		http.Error(res, "Existing user error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}

// login recieves a username and password and checks if those are a valid pair.
func login(res http.ResponseWriter, req *http.Request) {
	//logging
	req.ParseForm()
	username := html.EscapeString(req.FormValue("username"))
	password := html.EscapeString(req.FormValue("password"))
	log.Println(time.Now().Format(time.RFC850), "User Login Attempt by: ", username)
	var databaseUsername string
	var databasePassword string

	err := database.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login?retry=1", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login?retry=1", 301)
		return
	}

	res.Write([]byte("Hello " + databaseUsername))

}
