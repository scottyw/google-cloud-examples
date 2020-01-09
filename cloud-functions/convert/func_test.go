package convert

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFahrenheitToCelsius0(t *testing.T) {
	r := httptest.NewRequest("POST", "/", strings.NewReader("32"))
	w := httptest.NewRecorder()
	FahrenheitToCelsius(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "32°F is 0°C\n", string(body))
}

func TestFahrenheitToCelsius100(t *testing.T) {
	r := httptest.NewRequest("POST", "/", strings.NewReader("212"))
	w := httptest.NewRecorder()
	FahrenheitToCelsius(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "212°F is 100°C\n", string(body))
}

func TestFahrenheitToCelsiusMinus40(t *testing.T) {
	r := httptest.NewRequest("POST", "/", strings.NewReader("-40"))
	w := httptest.NewRecorder()
	FahrenheitToCelsius(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "-40°F is -40°C\n", string(body))
}

func TestFahrenheitToCelsiusBEmptyBody(t *testing.T) {
	r := httptest.NewRequest("POST", "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	FahrenheitToCelsius(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, "failed to convert body to an int\n", string(body))
}
