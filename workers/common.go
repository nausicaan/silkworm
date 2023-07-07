package workers

import (
	"os"
	"os/exec"
)

func scribe(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0666))
}

// Run a command with no verbosity
func execute(name string, task ...string) {
	lpath, err := exec.LookPath(name)
	inspect(err)
	exec.Command(lpath, task...).CombinedOutput()
}

// Run a command, then capture and return the output as a byte
func capture(name string, task ...string) []byte {
	lpath, err := exec.LookPath(name)
	inspect(err)
	osCmd, _ := exec.Command(lpath, task...).CombinedOutput()
	return osCmd
}

// Run standard commands and print the output to the terminal
func verbose(name string, task ...string) {
	path, err := exec.LookPath(name)
	osCmd := exec.Command(path, task...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err = osCmd.Run()
	inspect(err)
}

// Check for errors, log the result if found
func inspect(err error) {
	if err != nil {
		panic(err)
	}
}

// Remove files or directories
func cleanup(cut ...string) {
	inspect(os.Remove(cut[0.]))
}
