package main

// Launch the program and execute according to the supplied flag
func main() {
	if len(flag) < 1 {
		alert("No arguments detected -")
		about()
	} else {
		switch flag {
		case "-c", "--create":
			clearout()
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
