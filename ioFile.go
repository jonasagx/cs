package main

import . "github.com/jonasagx/csutils"

func main() {
	var text string
	text = ReadFileContent("/tmp/dat")
	Print(text)
}