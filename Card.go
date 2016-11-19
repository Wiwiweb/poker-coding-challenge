package main

type Card struct {
	Value CardValue
	Suit  string
}

// Having an enum-style type allows us to deal with Jacks and up
// without having to translate them to ints,
// which makes the code a bit clearer

type CardValue int

const (
	TWO CardValue = iota
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

func (c Card) Equals(otherCard Card) bool {
	return c.Value == otherCard.Value
}

func (c Card) IsBetterThan(otherCard Card) bool {
	return c.Value > otherCard.Value
}

// sort.Interface methods to allow sorting card values with builtin sort functions
type CardsByValue []CardValue

func (c CardsByValue) Len() int {
	return len(c)
}
func (c CardsByValue) Less(i, j int) bool {
	return c[i] < c[j]
}
func (c CardsByValue) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
