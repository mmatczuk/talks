package ssh

import (
	"io"
	"time"
)

type Client interface {
	Output(command string) (string, error)
	Shell(args ...string) error

	// Start starts the specified command without waiting for it to finish. You
	// have to call the Wait function for that.
	//
	// The first two io.ReadCloser are the standard output and the standard
	// error of the executing command respectively. The returned error follows
	// the same logic as in the exec.Cmd.Start function.
	Start(command string) (io.ReadCloser, io.ReadCloser, error) // HL

	// Wait waits for the command started by the Start function to exit. The
	// returned error follows the same logic as in the exec.Cmd.Wait function.
	Wait() error
}

type Cmd struct {
	// Command is the command to run remotely. This is executed as if
	// it were a shell command, so you are expected to do any shell escaping
	// necessary.
	Command string // HL
	// Stdin specifies the process's standard input. If Stdin is
	// nil, the process reads from an empty bytes.Buffer.
	Stdin io.Reader
	// Stdout and Stderr represent the process's standard output and
	// error.
	//
	// If either is nil, it will be set to ioutil.Discard.
	Stdout io.Writer
	Stderr io.Writer
	// Internal fields
}

// Communicator is an interface that must be implemented by all communicators
// used for any of the provisioners
type Communicator interface {
	// Connect is used to setup the connection
	Connect(terraform.UIOutput) error
	// Disconnect is used to terminate the connection
	Disconnect() error
	// Timeout returns the configured connection timeout
	Timeout() time.Duration
	// ScriptPath returns the configured script path
	ScriptPath() string
	// Start executes a remote command in a new session
	Start(*remote.Cmd) error // HL
	// Upload is used to upload a single file
	Upload(string, io.Reader) error // HL
	// UploadScript is used to upload a file as an executable script
	UploadScript(string, io.Reader) error // HL
	// UploadDir is used to upload a directory
	UploadDir(string, string) error // HL
}
