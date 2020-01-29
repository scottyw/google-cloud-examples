package convertbuilder

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func build(convert func(int64) int64) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}
		in, err := strconv.ParseInt(string(body), 10, 64)
		if err != nil {
			http.Error(w, "failed to convert body to an int", http.StatusInternalServerError)
			return
		}
		out := convert(in)
		fmt.Fprintf(w, "%d converts to %d\n", in, out)
	}
}

func fahrenheitToCelsius(in int64) int64 {
	return 5 * (in - 32) / 9
}

func celsiusToFahrenheit(in int64) int64 {
	return (9 * in / 5) + 32
}

func kelvinToCelsius(in int64) int64 {
	return in - 273
}

func celsiusToKelvin(in int64) int64 {
	return in + 273
}

// Run this command inside the "cloud-functions/convertbuilder" dir
// gcloud functions deploy --source=. --trigger-http --runtime=go111 --allow-unauthenticated KelvinToCelsius

// FahrenheitToCelsius converts F to C
var FahrenheitToCelsius = build(fahrenheitToCelsius)

// CelsiusToFahrenheit converts C to F
var CelsiusToFahrenheit = build(celsiusToFahrenheit)

// KelvinToCelsius converts K to C
var KelvinToCelsius = build(kelvinToCelsius)

// CelsiusToKelvin converts C to K
var CelsiusToKelvin = build(celsiusToKelvin)
