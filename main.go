package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func hellowHandeler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string `json:"message"`
	}
	responcedwithJson(w, http.StatusOK, response{Message: "Hello, World!"})
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Server is running on port %s\n", port)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Set up the router and CORS middleware
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r1 := chi.NewRouter()
	r1.HandleFunc("/hello", hellowHandeler)
	router.Mount("/api", r1)

	// Start the server
	err := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on http://localhost:%s\n", port)
	if err := err.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
