package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	uuidv4, err := uuid.NewRandom()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	_, err = fmt.Fprintln(w, uuidv4.String(), "")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	fmt.Printf("Listen Server Port: %v", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
