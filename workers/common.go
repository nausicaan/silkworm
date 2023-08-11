package workers

import (
	"os"
	"os/exec"
)

func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0666))
}

func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-e":
		exec.Command(task, args...).CombinedOutput()
	case "-c":
		both, _ := osCmd.CombinedOutput()
		return both
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
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
