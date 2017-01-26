package main

// Library used for this example
import . "github.com/jonasagx/csutils"

func main() {
	var name string = ReadName("Write your name then press ENTER, please: ")
	Print("The name", name, "has", len(name), "characters")

	var ageInDays int = ReadInteger("Write your age please, then press ENTER: ")
	Print("You've lived", ageInDays, "days")
}