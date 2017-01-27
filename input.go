package main

// Library used for this example
import . "github.com/jonasagx/csutils"

func main() {
	const daysPerYearOnEarth int = 365

	var name string
	var nameLength int

	name = ReadName("Write your name then press ENTER, please: ")
	nameLength = len(name)

	Print("The name", name, "has", nameLength, "characters")

	var ageInYears int = ReadInteger("Write your age please, then press ENTER: ")
	var ageInDays = ageInYears * daysPerYearOnEarth
	Print(name, "you've lived", ageInDays, "days")
}