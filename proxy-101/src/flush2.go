// flushAfterChunkWriter works with net/http/internal.chunkedWriter
// and forces a flush after each chunk is written.
// There is also net/http/internal.FlushAfterChunkWriter
// that does the same thing nicer, but it is not available.
type flushAfterChunkWriter struct {
	*bufio.Writer
}

func (w flushAfterChunkWriter) WriteString(s string) (n int, err error) {
	n, err = w.Writer.WriteString(s)
	if s == "\r\n" && err == nil {
		err = w.Flush()
	}
	return
}
