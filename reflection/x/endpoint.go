package x

import (
	"log"
	"net/http"
)

var debug bool = true

// START OMIT
func search(w http.ResponseWriter, r *http.Request) {
	// URL /search?sid=id&l=foo&l=bar&max=100&x=true
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data SearchRequest
	if err := bindParams(r, &data); err != nil { // HL
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("%+v", data) // "RequestContext:{SessionID:id} Labels:[foo bar] MaxResults:100 Exact:true}"
}

// END OMIT
