package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var isHeistOn bool = true
	var eludedGuards int = rand.Intn(100)

	if eludedGuards > 50 {
		fmt.Println("Looks like you've managed to make it past the gaurds.  Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	var openedVault int = rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The Vault couldn't be opened!")
	}

	var leftSafely int = rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("The alarm was tripped! Your heist has failed!")
		case 1:
			isHeistOn = false
			fmt.Println("You forgot your tools! Your heist has failed")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		var amtStolen int = 10000 + rand.Intn(1000000)
		fmt.Printf("YOu got away with $%d dollars!", amtStolen)
	}
}
