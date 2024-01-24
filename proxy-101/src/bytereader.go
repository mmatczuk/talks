type byteReader struct {
	r io.Reader
}

func (r byteReader) Read(p []byte) (int, error) {
	return r.r.Read(p[:1])
}
