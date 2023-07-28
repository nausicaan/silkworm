package workers

import (
	"os"
	"os/exec"
)

func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0666))
}

// Run a command with no verbosity
func execute(name string, task ...string) {
	lpath, err := exec.LookPath(name)
	inspect(err)
	exec.Command(lpath, task...).CombinedOutput()
}

// Run a command, then capture and return the output as a byte variable
func capture(name string, task ...string) []byte {
	lpath, err := exec.LookPath(name)
	inspect(err)
	osCmd, _ := exec.Command(lpath, task...).CombinedOutput()
	return osCmd
}

// Run a command, then print the output to the terminal
func verbose(name string, task ...string) {
	path, err := exec.LookPath(name)
	osCmd := exec.Command(path, task...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err = osCmd.Run()
	inspect(err)
}

// Check for errors, halt the program if found
func inspect(err error) {
	if err != nil {
		panic(err)
	}
}

// Remove files or directories
func cleanup(cut ...string) {
	inspect(os.Remove(cut[0.]))
}
