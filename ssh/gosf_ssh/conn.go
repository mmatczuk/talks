package gosf_ssh

import (
	"log"

	"golang.org/x/crypto/ssh"
)

// Connect dials an SSH Client connected to host, using the provided methods
// to authenticate
func Connect(host string, methods ...ssh.AuthMethod) (*ssh.Client, error) {
	cfg := ssh.ClientConfig{
		User: "chris",
		Auth: methods, // HL
	}

	return ssh.Dial("tcp", host, &cfg) // HL
}

// END Connect OMIT

// Run demonstrates dialing an SSH Client with multiple methods
func Run() { // OMIT
	agent, err := SSHAgent() // HL
	// handle error
	if err != nil { // OMIT
		log.Fatal(err) // OMIT
	} // OMIT

	keyPair, err := KeyPair("/home/chris/.ssh/id_rsa") // HL
	// handle error
	if err != nil { // OMIT
		log.Fatal(err) // OMIT
	} // OMIT

	client, err := Connect("example.com:22", agent, keyPair) // HL
	// handle error
	if err != nil { // OMIT
		log.Fatal(err) // OMIT
	} // OMIT

	defer client.Close() // HL
} // END Run OMIT
