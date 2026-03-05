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
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))
	r.HandleFunc("/hello", hellowHandeler)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}

}
