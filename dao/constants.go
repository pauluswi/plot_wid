package dao

import (
	"errors"
)

var (
	// ErrMaxSlotReached - Parking Full
	ErrMaxSlotReached = errors.New("Sorry, parking lot is full")
	// ErrCarNotFound - Car Not Found
	ErrCarNotFound = errors.New("Not found")
	// ErrInvalidMaxSlots - Reach Max Slot
	ErrInvalidMaxSlots = errors.New("Max slots should be greter than 0")
)
