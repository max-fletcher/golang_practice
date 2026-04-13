package main

// # IMPORTING PACKAGES:
// To import 3rd party modules for go, import and use the package in your file then use
// "go mod tidy"(cleanup unused deps) and then "go mod vendor"(copy only what’s needed).
// Although some packages like "github.com/google/uuid" and "github.com/lib/pq"
// requires the "go get" command to install.
// Here are some other useful commands:
// "go list -m all" to see a list of all packages
// "go mod why github.com/go-chi/chi/v5" to check why a package exists
// "go install github.com/some/tool@latest"; only use "get" when installing binaries
// Remember that this is not the same as downloading packages like "github.com/sqlc-dev/sqlc/cmd/sqlc@latest" or
// "go install github.com/pressly/goose/v3/cmd/goose@latest" that we have
// to use the command "go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest" to get. The reason being that sqlc is a tool
// (a set of binaries to be specific) that are not part of our application i.e nothing is downloaded into vendor and
// when we use "go build -o {filename}.exe", it is not compiled into the application.
import (
	"database/sql"
	"dummy/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Reference to a DB connection. Will be used to query data form DB.
type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Web Server made in GO")

	// We are using this because we need to set a bridge, When we use "os.Getenv" it only checks the os's
	// env variables and not the .env file. In golang .env is not automatically loaded into the os's
	// environment variables.
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in env")
	}

	// using go's sql package from its standard library to establish connection
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to the database", err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	fmt.Printf("PORT: %v", portString)

	router := chi.NewRouter()
	router.Use(httprate.Limit(
		100,           // number of requests
		1*time.Minute, // time window
		// The third parameter and onwards accepts multiple option functions(httprate.Limit is variadic funtion), which are handler functions
		// that have many variations. e.g "httprate.WithLimitHandler" is a custom handler function that you can add more logic
		// to(if you need to). Here, it sends back a custom json response. But if you want to choose which key to use to rate-limit,
		// you can use "httprate.WithKeyFuncs" where you can define what keys should be used to rate-limit requests like this:
		// httprate.WithKeyFuncs(func(r *http.Request) (string, error) {
		// 	return r.Header.Get("X-User-ID"), nil
		// }),
		// You can also pass multiple of these; each option function controlling an aspect of the rate-limiter e.g:
		// httprate.Limit(
		// 	100,
		// 	time.Minute,
		// 	httprate.WithKeyFuncs(httprate.KeyByIP),
		// 	httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
		// 		respondWithJSON(w, 429, map[string]string{
		// 			"error": "rate limit exceeded",
		// 		})
		// 	}),
		// )
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			respondWithJSON(w, 429, response{
				Code:    429,
				Status:  "ok",
				Message: "Too many requests. Rate limit exceeded.",
			})
		}),
	))

	// CORS
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })

	// Making a sub-router
	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", handlerError)
	v1router.Post("/users", apiCfg.handlerCreateUser) // route for creating users in DB
	router.Mount("/v1", v1router)

	// Server options like router and port
	// On windows, to start the server, use "go build -o {filename}.exe" then ".\go-rss-agg.exe"
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = server.ListenAndServe() // initialize server
	if err != nil {               // throws an error if the server fails
		log.Fatal(err)
	}
}
