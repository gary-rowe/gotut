package main

import "fmt"

const sixteenBitMax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {

	// Remember uppercase means public
	GasPedal      uint16 // min -, max 65535
	BrakePedal    uint16
	SteeringWheel uint16
	TopSpeedKmh   float64
}

// This is a value receiver function (no modification)
func (c car) mph() float64 {
	return float64(c.GasPedal) * (c.TopSpeedKmh / sixteenBitMax / kmhMultiple)
}

// This is a pointer receiver function (allows modification)
func (c *car) newTopSpeed(newSpeed float64) {
	c.TopSpeedKmh = newSpeed
}

func main() {
	aCar := car{
		GasPedal:      65000,
		BrakePedal:    0,
		SteeringWheel: 12561,
		TopSpeedKmh:   225.0}

	fmt.Println("Gas pedal:", aCar.GasPedal)

	// Use the value receiver
	fmt.Println("MPH:", aCar.mph())

	// Use the pointer receiver
	aCar.newTopSpeed(500)

	fmt.Println("New top speed:", aCar.TopSpeedKmh)
	fmt.Println("MPH:", aCar.mph())

}
