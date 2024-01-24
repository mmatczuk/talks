uconn, ok := res.Body.(io.ReadWriteCloser)
if !ok {
	log.Errorf(res.Request.Context(), "internal error: switching protocols response with non-writable body")
	return errClose
}

