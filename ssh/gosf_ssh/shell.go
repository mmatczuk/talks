package gosf_ssh

import (
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func RunShell(sess *ssh.Session) { // OMIT
	sess.Stdin = os.Stdin
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	modes := ssh.TerminalModes{ // HL
		ssh.ECHO:          1,      // please print what I type
		ssh.ECHOCTL:       0,      // please don't print control chars
		ssh.TTY_OP_ISPEED: 115200, // baud in
		ssh.TTY_OP_OSPEED: 115200, // baud out
	}

	termFD := int(os.Stdin.Fd())

	w, h, _ := terminal.GetSize(termFD) // HL

	termState, _ := terminal.MakeRaw(termFD) // HL
	defer terminal.Restore(termFD, termState)

	sess.RequestPty("xterm-256color", h, w, modes) // HL
	sess.Shell()                                   // HL
	sess.Wait()                                    // HL
} // OMIT
