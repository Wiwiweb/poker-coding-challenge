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
