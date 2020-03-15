package dao

// Car initiate Car Struct
type Car struct {
	colour string
	regNo  string
}

// GetRegNo - Initiation
func (c *Car) GetRegNo() string {
	return c.regNo
}

// GetColour - Initiation
func (c *Car) GetColour() string {
	return c.colour
}

// NewCar - Initiation
func NewCar(regNo, colour string) *Car {
	return &Car{
		regNo:  regNo,
		colour: colour,
	}
}
