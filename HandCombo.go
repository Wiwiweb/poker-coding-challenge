package main

type HandCombo struct {
	value          ComboValue
	isThisCombo    func([5]Card) bool
	getTiebreakers func([5]Card) []CardValue
}

type ComboValue int

const (
	HIGH_CARD ComboValue = iota
	PAIR
	THREE_OF_A_KIND
	FLUSH
)

func (f *HandCombo) isBetterThan(otherCombo HandCombo) bool {
	return f.value > otherCombo.value
}
