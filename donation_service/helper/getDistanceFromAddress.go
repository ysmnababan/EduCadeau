package helper

import (
	"fmt"
	"math"
)

func CalculateShippingCost(distanceMeters int) int {
	// Define constants
	const initialCostPerMeter = 15   // Initial cost per meter
	const baseDistance = 20000       // Base distance after which dynamic pricing starts
	const fixedCostBelowBase = 20000 // Fixed cost for distances less than baseDistance

	// Calculate shipping cost
	if distanceMeters < baseDistance {
		return fixedCostBelowBase
	} else {
		const decreaseFactor = 1.7

		// Calculate cost
		distanceBeyondBase := distanceMeters - baseDistance
		costReduction := decreaseFactor * math.Log(float64(distanceBeyondBase)+1) // Adding 1 to avoid log(0)
		costPerMeter := initialCostPerMeter - costReduction

		// Ensure cost per meter not minus
		if costPerMeter < 0 {
			costPerMeter = 0.25
		}

		totalCost := fixedCostBelowBase + int(math.Ceil(float64(distanceBeyondBase)*costPerMeter))
		// Round up to nearest 1000
		roundedCost := int(math.Ceil(float64(totalCost)/1000.0)) * 1000
		fmt.Println("dist : ", distanceBeyondBase)
		fmt.Println("cpm : ", costPerMeter)
		fmt.Println("creduct : ", costReduction)
		return roundedCost/10

	}
}
