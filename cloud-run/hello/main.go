package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env var must be specified")
	}
	http.HandleFunc("/", handler)
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Listening on %s", addr)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello: " + os.Getenv("K_REVISION")))
}
