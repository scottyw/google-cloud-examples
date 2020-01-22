package convert

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// Foo spews the request into the response
func Foo(w http.ResponseWriter, r *http.Request) {
	spew.Fdump(w, r)
}
