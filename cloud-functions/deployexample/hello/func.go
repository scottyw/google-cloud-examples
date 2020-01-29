package hello

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/google-cloud-examples/cloud-functions/deployexample/common"
)

// Run this command inside the "cloud-functions/deployexample" dir
// gcloud functions deploy --source=. --trigger-http --runtime=go111 --allow-unauthenticated Hello

// Hello pretty prints the request into the response
func Hello(w http.ResponseWriter, r *http.Request) {
	common.Log("hello")
	spew.Fdump(w, r)
}
