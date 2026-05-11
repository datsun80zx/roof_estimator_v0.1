package entities

type RoofDimensions struct {
	Area         float32
	RidgeLength  float32
	HipLength    float32
	EaveLength   float32
	RakeLength   float32
	ValleyLength float32
	Perimeter    float32
}

type RoofDetails struct {
	IntakeVentStyle   string
	ExhaustVentStyle  string
	NumberOfSkylights int
	NumberOfVents     int
	NumberOfChimney   int
	NumberOfBroan     int
	NumberOfCrickets  int
	NumberOfLayers    int
	DeckingMaterial   int
	NumberOfDecking   int
	WasteFactor       int
}

type RoofingMaterialList struct {
	Decking int
}

/*
plank decking std dimension: 10' sections
plywood std dimension: 4'x 8' sheets
osb std dimensions: 4' x 8' sheets
*/

type PlankDeck struct {
	length   int
	quantity int
}

type OSB struct {
	area     int
	quantity int
}

type Plywood struct {
	area     int
	quantity int
}

type DripEdge struct {
	style    string
	length   int
	quantity int
	color    string
}

type IceWaterShield struct {
	brand    string
	style    string
	lenght   int
	quantity int
}

type Underlayment struct {
	brand    string
	style    string
	length   int
	quantity int
}

type StarterStrip struct {
	brand    string
	style    string
	length   int
	quantity int
}

type Shingles struct {
	brand    string
	style    string
	length   int
	quantity int
	color    string
}

type HipRidgeShingles struct {
	brand    string
	style    string
	length   int
	quantity int
	color    string
}

type ExhaustVent struct {
	style    string
	quantity int
	length   int
}

type IntakeVent struct {
	style    string
	quantity int
	length   int
}

type ApronFlashing struct {
	color    string
	quantity int
	lenght   int
}

type StepFlashing struct {
	color    string
	quantity int
}

type CounterFlashing struct {
	color    string
	quantity int
}

type ValleyFlashing struct {
	color  string
	length int
}
