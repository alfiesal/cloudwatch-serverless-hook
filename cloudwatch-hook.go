package main

import (
	"bytes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	data := []byte(`{"First Name":"Alfie","Last Name":"Salomons"}`)

	url := os.Getenv("URL")
	_, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}
}
