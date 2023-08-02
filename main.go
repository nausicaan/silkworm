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
	zero   string = "Insufficient arguments supplied -"
)

// Launch the program and execute according to the supplied flag
func main() {
	switch os.Args[1] {
	case "-v", "--version":
		fmt.Println(yellow+"Silkworm", green+bv)
		fmt.Println(reset)
	case "-h", "--help":
		about()
	default:
		workers.Quarterback()
	}
}

// about prints help information for using the program
func about() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [vendor/plugin]:[version]")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Against your composer.json file, run:")
	fmt.Println(green + "    silkworm wpackagist-plugin/mailpoet:4.6.1")
	fmt.Println(yellow, "\nAdditional Options:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -v, --version", reset, "	Display App Version")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/nausicaan/silkworm.git")
	fmt.Println(reset)
}
