package x

import (
	"fmt"
	"net/http"
)

var debug bool = true

// START OMIT
func search(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// handle error
		return
	}
	var data SearchRequest
	if err := unpack(r, &data); err != nil { // HL
		// handle error
		return
	}
	if debug {
		fmt.Printf("%+v", data) // HL
	}
}

// END OMIT
