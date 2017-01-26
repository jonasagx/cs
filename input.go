package main

// Library used for this example
import . "github.com/jonasagx/csutils"

func main() {
	const daysPerYear int = 365

	var name string = ReadName("Write your name then press ENTER, please: ")
	Print("The name", name, "has", len(name), "characters")

	var ageInYears int = ReadInteger("Write your age please, then press ENTER: ")
	var ageInDays = ageInYears * daysPerYear
	Print(name, "you've lived", ageInDays, "days")
}