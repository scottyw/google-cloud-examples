package main

import (
	"fmt"
	"net/http"

	"github.com/google-cloud-examples/cloud-functions/convert"
)

func main() {
	fmt.Println("Listening on http://localhost:8080")
	http.HandleFunc("/", convert.FahrenheitToCelsius)
	http.ListenAndServe(":8080", nil)
}
