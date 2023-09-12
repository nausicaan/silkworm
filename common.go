package main

import (
	"fmt"
	"os"
	"os/exec"
)

var passed = os.Args

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0666))
}

// Run a terminal command using flags to customize the output
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

// about prints help information for using the program
func about() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag] [vendor/plugin]:[version]")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Adding your path to file if necessary, run:")
	fmt.Println(green + "    silkworm -c wpackagist-plugin/mailpoet:4.6.1")
	fmt.Println(yellow, "\nAdditional Options:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -v, --version", reset, "	Display App Version")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/nausicaan/silkworm.git")
	fmt.Println(reset)
}

// Alert prints a colourized error message
func alert(message string) {
	fmt.Println("\n"+red, message, halt, reset)
}

// Display the build version of the program
func build() {
	fmt.Println(yellow+"Silkworm", green+bv, reset)
}
