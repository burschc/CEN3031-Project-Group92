package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"ufpmp/httpd"

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

// DatabaseName is the filename of the Database which stores user account information.
const DatabaseName = "accountDB"

type Close func()
type DB struct {
	Name     string
	Database *sql.DB
	Close
}

var Database DB

func DatabaseHandlers(r *mux.Router) {
	r.HandleFunc("/signup", signup).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/login", login).Methods(http.MethodPost, http.MethodOptions)

	log.Print("Registered Database handlers.")
}

func DeclareDatabase(dbname string) {
	//Open the Database.
	Database.Database, _ = sql.Open("sqlite3", dbname)
	Database.Name = dbname
	Database.Close = func() {
		err := Database.Database.Close()
		if err != nil {
			log.Printf("Unable to close Database '%v'", Database.Name)
		}
	}
	//Make the Database if it doesn't exist.
	statement, _ := Database.Database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, passtype INTEGER)")
	_, _ = statement.Exec()
}

func LegacySetupOrOpenBasicDatabase() {
	//Open the Database.
	Database.Database, _ = sql.Open("sqlite3", "testdb.db")
	//Make the Database if it doesn't exist.
	statement, _ := Database.Database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT, passtype INTEGER)")
	_, _ = statement.Exec()
	//Prepare the to add a user to the Database, then actually do it.
	statement, _ = Database.Database.Prepare("INSERT INTO users (username, password, passtype) VALUES (?, ?, ?)")
	_, _ = statement.Exec("TestUser", "TestPassword", 1)
	//Prepare to print the entire list of entries.
	rows, _ := Database.Database.Query("SELECT id, username, password, passtype FROM users")
	var id int
	var username string
	var password string
	var passtype int
	for rows.Next() {
		rows.Scan(&id, &username, &password, &passtype)
		fmt.Println(strconv.Itoa(id) + ": " + username + " " + password + " " + strconv.Itoa(passtype))
	}
}

// printDatabase is a test function which prints the contents of the Database.
func printDatabase() {
	//Prepare to print the entire list of entries.
	rows, _ := Database.Database.Query("SELECT id, username, password, passtype FROM users")
	var id int
	var username string
	var password string
	var passtype int
	for rows.Next() {
		rows.Scan(&id, &username, &password, &passtype)
		fmt.Println(strconv.Itoa(id) + ": " + username + " | " + password + " | " + strconv.Itoa(passtype))
	}
}

// signup adds a specified account with given username and password to the Database if it is not already present.
func signup(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		httpd.PipeError(res, err)
	}

	username := html.EscapeString(req.FormValue("username"))
	password := html.EscapeString(req.FormValue("password"))

	var user string

	err := Database.Database.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err2 != nil {
			http.Error(res, "Hash error, unable to create your account.", http.StatusInternalServerError)
			return
		}
		//log.Printf(string(hashedPassword))
		//passtype -1 being "no passtype declared" -> defaulting to showing all types
		http.Error(res, "entered username is "+username+" and password was "+password, http.StatusOK)
		statement, _ := Database.Database.Prepare("INSERT INTO users (username, password, passtype) VALUES (?, ?, ?)")
		_, err2 = statement.Exec(username, hashedPassword, -1)
		if err2 != nil {
			http.Error(res, "Insert error, unable to create your account.", http.StatusInternalServerError)
			return
		}
		//if _, err = Database.Exec("INSERT INTO users (username, password, passtype) VALUES(?, ?, ?)", username, hashedPassword, -1); err != nil {
		//	http.Error(res, "Insert error, unable to create your account.", 500)
		//	return
		//}

		res.Write([]byte("User created!"))
	case err != nil:
		http.Error(res, "Existing user error, unable to create your account.", http.StatusConflict)
	default:
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Unknown error."))
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

	err := Database.Database.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("No user found with that username."))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("No user found with that password."))
		return
	}

	res.Write([]byte("Hello " + databaseUsername))

}
