package gosf_ssh

import (
	"log"

	"os"

	"golang.org/x/crypto/ssh"
)

func RunCmd(client *ssh.Client) { // OMIT
	sess, err := client.NewSession() // HL
	// handle error
	if err != nil { // OMIT
		log.Fatal(err) // OMIT
	} // OMIT
	defer sess.Close() // HL

	sess.Stdout = os.Stdout                          // HL
	sess.Setenv("LS_COLORS", os.Getenv("LS_COLORS")) // HL

	err = sess.Run("ls -lah") // HL
	// handle error
	if err != nil { // OMIT
		log.Fatal(err) // OMIT
	} // OMIT
} // OMIT
