package main

import (
	"fmt"
	"os"

	"github.com/nausicaan/silkworm/workers"
)

const (
	bv     string = "1.0"
	reset  string = "\033[0m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[41m"
	halt   string = "program halted"
)

// Launch the program and execute according to the supplied flag
func main() {
	if len(os.Args) == 1 {
		alert("No arguments detected -")
		about()
	} else {
		switch os.Args[1] {
		case "-c", "--create":
			workers.Quarterback()
		case "-h", "--help":
			about()
		case "-v", "--version":
			fmt.Println(yellow+"Silkworm", green+bv)
			fmt.Println(reset)
		default:
			alert("Unknown argument(s) supplied -")
			about()
		}
	}
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
