package deployexample

import (
	"github.com/google-cloud-examples/cloud-functions/deployexample/goodbye"
	"github.com/google-cloud-examples/cloud-functions/deployexample/hello"
)

// Hello is a reference to allow deployment in GCF
var Hello = hello.Hello

// Goodbye is a reference to allow deployment in GCF
var Goodbye = goodbye.Goodbye
