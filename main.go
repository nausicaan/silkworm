package main

// Launch the program and execute according to the supplied flag
func main() {
	if len(flag) == 2 {
		switch flag[1] {
		case "-h", "--help":
			about()
		case "-v", "--version":
			build()
		default:
			alert("Unknown argument(s) supplied -")
			about()
		}
	} else {
		clearout(common + "temp/")
		message("Creating tickets")
		serialize()
		sifter()
	}
}
