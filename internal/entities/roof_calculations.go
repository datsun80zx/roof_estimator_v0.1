package entities

/* all values are in feet unless otherwise specified */

var dripEdgeStick int = 10
var iceWaterShieldRoll int = 66
var underlaymentRoll int = 10 // this is in square feet
var starterStripBundle int = 105
var shingleBundle int = 33         // this is in square feet
var hipRidgeShingleBundle int = 33 // this is in linear feet
var ridgeVentSection int = 4
var inflowVentSection int = 4

func countOfDripStick(roof RoofDimensions) float32 {
	return float32(roof.Perimeter) / float32(dripEdgeStick)
}

func countOfIWRolls(roof RoofDimensions) float32 {
	return float32(roof.ValleyLength + (roof.EaveLength * 2) + roof.RidgeLength)
}
