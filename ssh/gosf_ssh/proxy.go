package gosf_ssh

import "golang.org/x/crypto/ssh"

func Proxy(bastion *ssh.Client, host string, clientCfg *ssh.ClientConfig) *ssh.Client {
	netConn, _ := bastion.Dial("tcp", host) // HL

	conn, chans, reqs, _ := ssh.NewClientConn(netConn, host, clientCfg) // HL

	return ssh.NewClient(conn, chans, reqs) // HL
}
