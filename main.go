package main

// Launch the program and execute according to the supplied flag
func main() {
	if inputs == 1 {
		alert("No arguments detected -")
		about()
	} else {
		switch passed[1] {
		case "-c", "--create":
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
}
