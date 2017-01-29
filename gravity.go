package main

// Author: Jack Wagner

// Library used for this example
import(
	. "github.com/jonasagx/csutils"
	"math"
)

func conversionLbsToKg(weightLbs int) float64{
	const conversion float64 = 2.2
	weightkg := float64(weightLbs)/conversion
	return weightkg
}

func main() {
	const g float64 = 9.8
	
	var weightLbs int
	var weightkg float64
	var mass float64
	
	weightLbs = ReadInteger("What is your weight? (in Lbs): ")
	weightkg = conversionLbsToKg(weightLbs)
	mass = float64(weightkg)/g
	Print("Your mass is:", mass)

	var decimals float64
	decimals = math.Remainder(mass, 1)
	//definition of rounding
	if decimals > 0.4 {
		mass=mass-decimals
		mass=mass+1
	}
	if decimals<0.5{
		mass=mass-decimals
	}
	//end result
	Print("Your rounded mass is:", mass)
}