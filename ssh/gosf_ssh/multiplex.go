package gosf_ssh

import (
	"bufio"

	"fmt"

	"log"

	"golang.org/x/crypto/ssh"
)

func TailLog(name string, client *ssh.Client, lines chan<- string) {
	sess, _ := client.NewSession()
	defer sess.Close()

	out, _ := sess.StdoutPipe() // HL

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)

	sess.Start("tail -f /var/log/app.log") // HL

	for scanner.Scan() {
		lines <- fmt.Sprintf("[%s] %s", name, scanner.Text())
	}

	sess.Wait() // HL
}

// END LIST OMIT

func MultiTail(bastion *ssh.Client, hosts []string, cfg *ssh.ClientConfig) {
	lines := make(chan string)

	for _, remote := range hosts {
		go TailLog( // HL
			remote,
			Proxy(bastion, remote, cfg), // HL
			lines,
		)
	}

	for l := range lines {
		log.Print(l)
	}
}
