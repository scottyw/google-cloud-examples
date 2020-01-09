package convert

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Run this command inside the "cloud-functions" dir
// gcloud functions deploy --source=convert --trigger-http --runtime=go111 --allow-unauthenticated FahrenheitToCelsius

// FahrenheitToCelsius converts celsius to fahrenheit
func FahrenheitToCelsius(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}
	c, err := strconv.ParseInt(string(body), 10, 64)
	if err != nil {
		http.Error(w, "failed to convert body to an int", http.StatusInternalServerError)
		return
	}
	f := 5 * (c - 32) / 9
	fmt.Fprintf(w, "%d°F is %d°C\n", c, f)
}
