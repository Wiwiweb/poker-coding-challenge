package main

type Card struct {
	value string
	suit  string
}

var valueOrder = []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

func (c Card) Equals(otherCard Card) bool {
	return c.value == otherCard.value
}

func (c Card) IsBetterThan(otherCard Card) bool {
	if c.Equals(otherCard) {
		return false
	}
	for _, value := range valueOrder {
		if c.value == value {
			return true
		}
		if otherCard.value == value {
			return false
		}
	}
	// We assume no invalid values, i.e. the value will be in the valueOrder slice
	// So this shouldn't ever reach this point
	return false
}
