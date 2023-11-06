package main

import "os"

// Launch the program and execute according to the supplied flag
func main() {
	if len(os.Args) < 1 {
		alert("No arguments detected -")
		about()
	} else {
		switch os.Args[1] {
		case "-c", "--create":
			// scopy()
			message("Creating tickets")
			serialize()
			sifter()
		case "-h", "--help":
			about()
		case "-v", "--version":
			build()
		default:
			alert("Unknown argument(s) supplied -")
			about()
		}
	}
	for _, v := range temp {
		cleanup(v)
	}
}
