package main

// # IMPORTING PACKAGES:
// To import 3rd party modules for go, import and use the package in your file then use
// "go mod tidy"(cleanup unused deps) and then "go mod vendor"(copy only what’s needed).
// Here are some other useful commands:
// "go list -m all" to see a list of all packages
// "go mod why github.com/go-chi/chi/v5" to check why a package exists
// "go install github.com/some/tool@latest"; only use "get" when installing binaries
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

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

	fmt.Printf("PORT: %v", portString)

	router := chi.NewRouter()

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
	router.Mount("/v1", v1router)

	// Server options like router and port
	// On windows, to start the server, use "go build -o {filename}.exe" then ".\go-rss-agg.exe"
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe() // initialize server
	if err != nil {             // throws an error if the server fails
		log.Fatal(err)
	}
}
