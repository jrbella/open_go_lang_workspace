package main

import "fmt"

// Create the function fuelGauge() here
func fuelGague(fuel int) {
	fmt.Printf("Pilot, you have %d gallons fuel left", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel

}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Printf("Hello! You have arrived at %v", planet)
	//forcing new line
	fmt.Println(" ")

}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining int
	var fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining = fuelRemaining - fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	// Test your functions!
	// Create `planetChoice` and `fuel`
	fmt.Println("Running script")
	var fuel int = 1000000
	var planetChoice = "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGague(fuel)
	// And then liftoff!

}
