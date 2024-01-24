func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	c := http.NewResponseController(rw)

	if err := c.Flush(); err != nil {
		// handle error
	}

	// write response body
}
