package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3030", router))
}
