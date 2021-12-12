package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func authCallback(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, r.URL)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func runApp() {
	port := os.Getenv("PORT")

	log.Printf("Running on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequests() {
	http.HandleFunc("/callback", authCallback)
}

func loadDotenv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadDotenv()

	handleRequests()

	runApp()
}
