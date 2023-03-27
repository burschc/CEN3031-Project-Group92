package mux_functions

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
	"ufpmp/httpd/app_handlers"
	"ufpmp/httpd/app_handlers/decal_filter"
)

// AppServerProperties is a struct which contains the properties used in the creation of the web app backend server.
type AppServerProperties struct {
	Router           *mux.Router           // Router is the GorillaMux router for the web app.
	Server           *http.Server          // Server is the http server for the web app.
	BackendURL       string                // BackendURL is the base URL of the backend http server. This is typically 'localhost'.
	BackendPort      string                // BackendPort is the port of the backend http server. This is typically '8080'.
	ContingentRoutes []func(r *mux.Router) // ContingentRoutes is an array of handler registration functions which depend on flags.
}

// ServerProperties is an implementation of AppServerProperties applicable to the web app.
var ServerProperties = AppServerProperties{
	BackendURL:  "localhost:",
	BackendPort: "8080",
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//											PRIVATE UTILITY FUNCTIONS												  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// corsHeaderMiddleware prepends a default CORS header that will work pretty much anywhere. For browser security, this
// really shouldn't be used in production.
func corsHeaderMiddleware(r *mux.Router) http.Handler {
	return handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
			"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
			"Cache-Control", "Content-Range", "Range"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
			"X-Requested-With", "If-Modified-Since", "Cache-Control",
			"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
		handlers.MaxAge(86400),
	)(r))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//											PUBLIC UTILITY FUNCTIONS												  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CreateAppServer creates the gorilla mux_functions router used by the web app, and the http server that uses that mux_functions router
// with all middlewares included.
func CreateAppServer(properties AppServerProperties) {
	//Create a new mux_functions router.
	ServerProperties.Router = mux.NewRouter()

	//Use the Custom-made CORS header middleware tied to the logging middleware for the server.
	rCustom := corsHeaderMiddleware(ServerProperties.Router)

	//Create a server with the following properties:
	ServerProperties.Server = &http.Server{
		Handler: rCustom,
		Addr:    properties.BackendURL + properties.BackendPort,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// RegisterHandlers registers all functions for the entire web application. It logs a message confirming that all paths
// in the function have been registered. Handlers should be registered in an order where the default is registered last.
func RegisterHandlers(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	decal_filter.DecalFilterHandlers(api)
	app_handlers.HttpHandlers(api)

	for _, f := range ServerProperties.ContingentRoutes {
		f(api)
	}

	app_handlers.DefaultHttpHandler(r)
}
