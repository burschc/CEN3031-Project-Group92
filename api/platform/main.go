package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"ufpmp/database"
	"ufpmp/httpd/mux_functions"
	"ufpmp/platform/arguments"
	"ufpmp/python"
)

func main() {

	//Process any command line arguments.
	arguments.ProcessArguments()

	//Set up the database, or open it if it already exists.
	database.DeclareDatabase(database.DatabaseName)

	//Make sure Python is installed and the virtual environment with all the required scripts is available.
	log.Print("Checking Python Virtual Environment...\n\n")
	python.SetupPythonVenv()

	//Create the router and server for the web app and register the handlers for the mux_functions.
	mux_functions.CreateAppServer(mux_functions.ServerProperties)
	mux_functions.RegisterHandlers(mux_functions.ServerProperties.Router)

	log.Printf("Starting http server accessible through %v...", mux_functions.ServerProperties.Server.Addr)

	go func() {
		if err := mux_functions.ServerProperties.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server encountered an error:\n%v", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ServerShutdown()
	os.Exit(0)
}

func ServerShutdown() {
	log.Print("Closing database...")
	database.Database.Close()
	log.Print("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), arguments.GracefulWait)
	defer cancel()

	if err := mux_functions.ServerProperties.Server.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Server has shut down successfully.")
}
