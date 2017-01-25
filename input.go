package main

// Library used for this example
import . "github.com/jonasagx/csutils"

func main() {
	var name string = ReadInput("Write your name then press ENTER, please: ")
	Print(name)
}