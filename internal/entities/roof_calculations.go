package entities

/* all values are in feet unless otherwise specified */

var dripEdgeStick int = 10
var iceWaterShieldRoll int = 66
var underlaymentRoll int = 1000 // this is in square feet
var starterStripBundle int = 105
var shingleBundle int = 33         // this is in square feet
var hipRidgeShingleBundle int = 33 // this is in linear feet
var ridgeVentSection int = 4
var inflowVentSection int = 4
var valleyFlashing int = 10

func CountOfDripStick(roof RoofDimensions) float32 {
	return float32(roof.Perimeter) / float32(dripEdgeStick)
}

func CountOfIWRolls(roof RoofDimensions) float32 {
	return (float32(roof.ValleyLength) + (float32(roof.EaveLength) * 2) + float32(roof.IntakeVent.Length)) / float32(iceWaterShieldRoll)
}

func CountOfUnderlayment(roof RoofDimensions) float32 {
	return float32(roof.Area) / float32(underlaymentRoll)
}

func CountOfStarter(roof RoofDimensions) float32 {
	return float32(roof.Perimeter) / float32(starterStripBundle)
}

func CountOfShingleBundle(roof RoofDimensions) float32 {
	return float32(roof.Area) / float32(shingleBundle)
}

func CountOfHipRidgeCap(roof RoofDimensions) float32 {
	return (float32(roof.HipLength) + float32(roof.RidgeLength)) / float32(hipRidgeShingleBundle)
}

func CountOfRidgeVent(roof RoofDimensions) float32 {
	return float32(roof.RidgeLength) / float32(ridgeVentSection)
}

func CountOfInflowVent(roof RoofDimensions) float32 {
	return float32(roof.IntakeVent.Length) / float32(inflowVentSection)
}

func CountOfValleyFlash(roof RoofDimensions) float32 {
	return float32(roof.ValleyLength) / float32(valleyFlashing)
}
