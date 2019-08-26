package gosf_ssh

import (
	"net"

	"net/http"

	"fmt"

	"golang.org/x/crypto/ssh"
)

func Tunnel(client *ssh.Client, localHost, remoteHost string) {
	listener, _ := net.Listen("tcp", localHost) // HL
	defer listener.Close()

	for {
		localConn, _ := listener.Accept()               // HL
		remoteConn, _ := client.Dial("tcp", remoteHost) // HL

		go copy(localConn, remoteConn)
		go copy(remoteConn, localConn)
	}
}

// END TUNNEL OMIT

func ReverseTunnel(client *ssh.Client, remoteHost string) {
	listener, _ := client.Listen("tcp", remoteHost) // HL
	defer listener.Close()

	handler := func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello, GoSF!")
	}

	http.Serve(listener, http.HandlerFunc(handler)) // HL
}
