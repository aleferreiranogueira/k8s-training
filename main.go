package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	fmt.Printf("Init API")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", list)
	log.Fatal(http.ListenAndServe(":8080", router))
}
