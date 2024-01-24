if bufReader != nil {
	// snippet borrowed from `proxy` plugin
	if n := bufReader.Reader.Buffered(); n > 0 {
		rbuf, err := bufReader.Reader.Peek(n)
		if err != nil {
			return http.StatusBadGateway, err
		}
		targetConn.Write(rbuf)
	}
}
