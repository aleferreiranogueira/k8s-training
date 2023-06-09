package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var up bool

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("API_HOST") == "" {
		fmt.Printf("API_HOST is not set! App is not ready\n")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	fmt.Printf("API_HOST set! App is READY!\n")
	w.WriteHeader(http.StatusOK)
}

func main() {
	up = false
	fmt.Printf("Initializing...\n")

	timeout := time.After(10 * time.Second)

	select {
	case <-timeout:
		fmt.Printf("App is Up!\n")
		up = true
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", list)
	router.HandleFunc("/health", healthCheck)
	log.Fatal(http.ListenAndServe(":8080", router))
}
