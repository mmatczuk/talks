tlsconn := tls.Server(&peekedConn{
	conn,
	io.MultiReader(bytes.NewReader(buf), conn),
}, p.mitm.TLSForHost(req.Host))
