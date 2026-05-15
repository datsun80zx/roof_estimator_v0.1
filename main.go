package main

import (
	"fmt"

	"github.com/datsun80zx/roof_estimator_v0.1/internal/entities"
)

func main() {
	fmt.Println("estimating roof...")

	intakeVent := entities.IntakeVent{
		Style:    "existing",
		Quantity: 0,
		Length:   0,
	}

	r1 := entities.RoofDimensions{
		Area:         6598.0,
		RidgeLength:  92.0,
		HipLength:    315.0,
		EaveLength:   455.0,
		RakeLength:   87.0,
		ValleyLength: 82.0,
		Perimeter:    542.0,
		IntakeVent:   intakeVent,
	}

	countOfShingles := entities.CountOfShingleBundle(r1)
	countOfUnderlayment := entities.CountOfUnderlayment(r1)
	countOfIW := entities.CountOfIWRolls(r1)
	countOfRidgeVent := entities.CountOfRidgeVent(r1)
	countOfCap := entities.CountOfHipRidgeCap(r1)
	countOfDrip := entities.CountOfDripStick(r1)
	countOfStarter := entities.CountOfStarter(r1)

	fmt.Printf("Total Counts: \n\nShingles: %v\nUnderlayment: %v\nIce & Water: %v\nRidge Vent: %v\nHip & Ridge Shingles: %v\nDrip Edge: %v\nStarter Shingles: %v",
		countOfShingles,
		countOfUnderlayment,
		countOfIW,
		countOfRidgeVent,
		countOfCap,
		countOfDrip,
		countOfStarter)
}
