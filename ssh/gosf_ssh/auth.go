package gosf_ssh

import (
	"io/ioutil"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// KeyPair creates an AuthMethod from a private key PEM file
func KeyPair(keyFile string) (ssh.AuthMethod, error) {
	pem, err := ioutil.ReadFile(keyFile) // HL
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(pem) // HL
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil // HL
}

// END KeyPair OMIT

// SSHAgent creates an AuthMethod leveraging keys stored in ssh-agent/pageant
func SSHAgent() (ssh.AuthMethod, error) {
	agentSock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")) // HL
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeysCallback(agent.NewClient(agentSock).Signers), nil // HL
}

// END SSHAgent OMIT
