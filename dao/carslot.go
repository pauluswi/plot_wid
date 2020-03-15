package dao

// CarSlot Initiation
type CarSlot struct {
	no      int
	vehicle Vehicle
}

// NewSlot - Initiation
func NewSlot(no int) *CarSlot {
	s := new(CarSlot)
	s.no = no
	return s
}

//SetNo - Set Car Slot
func (s *CarSlot) SetNo(no int) {
	s.no = no
}

//GetNo - Get Slot Number
func (s *CarSlot) GetNo() int {
	return s.no
}

//SetVehicle - Set Vehicle Slot
func (s *CarSlot) SetVehicle(v Vehicle) {
	s.vehicle = v
}

//GetVehicle - Get Vehicle Slot
func (s *CarSlot) GetVehicle() Vehicle {
	return s.vehicle
}
