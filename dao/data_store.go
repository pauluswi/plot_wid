package dao

// DataStore - Interface
type DataStore interface {
	Park(Vehicle) (Slot, error)
	Leave(Slot) error
	GetAll() ([]Slot, error)
	GetAllSlotsByColour(string) ([]Slot, error)
	GetSlotByRegNo(string) (Slot, error)
}
