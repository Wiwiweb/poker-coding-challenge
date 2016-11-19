package main

import (
	"testing"

	test_assert "github.com/stretchr/testify/assert"
)

func TestCalculateHandValue_Flush(t *testing.T) {
	assert := test_assert.New(t)

	cards := [5]Card{
		{Value: TWO, Suit: "H"},
		{Value: FIVE, Suit: "H"},
		{Value: SIX, Suit: "H"},
		{Value: JACK, Suit: "H"},
		{Value: THREE, Suit: "H"}}
	hand := Hand{Cards: cards}
	hand.calculateHandValue()

	assert.Equal(FLUSH, hand.combo.value)
	assert.Equal([]CardValue{JACK, SIX, FIVE, THREE, TWO}, hand.tiebreakerCardsValue)
}

func TestCalculateHandValue_ThreeOfAKind(t *testing.T) {
	assert := test_assert.New(t)

	cards := [5]Card{
		{Value: TWO, Suit: "H"},
		{Value: FIVE, Suit: "C"},
		{Value: JACK, Suit: "H"},
		{Value: FIVE, Suit: "H"},
		{Value: FIVE, Suit: "S"}}
	hand := Hand{Cards: cards}
	hand.calculateHandValue()

	assert.Equal(THREE_OF_A_KIND, hand.combo.value)
	assert.Equal([]CardValue{FIVE}, hand.tiebreakerCardsValue)
}

func TestCalculateHandValue_Pair(t *testing.T) {
	assert := test_assert.New(t)

	cards := [5]Card{
		{Value: TWO, Suit: "H"},
		{Value: FIVE, Suit: "C"},
		{Value: JACK, Suit: "H"},
		{Value: KING, Suit: "H"},
		{Value: FIVE, Suit: "S"}}
	hand := Hand{Cards: cards}
	hand.calculateHandValue()

	assert.Equal(PAIR, hand.combo.value)
	assert.Equal([]CardValue{FIVE, KING, JACK, TWO}, hand.tiebreakerCardsValue)
}

func TestCalculateHandValue_PairWithAnotherHigherPair(t *testing.T) {
	assert := test_assert.New(t)

	cards := [5]Card{
		{Value: TWO, Suit: "H"},
		{Value: FIVE, Suit: "C"},
		{Value: FIVE, Suit: "H"},
		{Value: JACK, Suit: "H"},
		{Value: JACK, Suit: "S"}}
	hand := Hand{Cards: cards}
	hand.calculateHandValue()

	assert.Equal(PAIR, hand.combo.value)
	assert.Equal([]CardValue{JACK, FIVE, FIVE, TWO}, hand.tiebreakerCardsValue)
}

func TestCalculateHandValue_HighCard(t *testing.T) {
	assert := test_assert.New(t)

	cards := [5]Card{
		{Value: TWO, Suit: "H"},
		{Value: FIVE, Suit: "C"},
		{Value: JACK, Suit: "H"},
		{Value: ACE, Suit: "H"},
		{Value: SIX, Suit: "S"}}
	hand := Hand{Cards: cards}
	hand.calculateHandValue()

	assert.Equal(HIGH_CARD, hand.combo.value)
	assert.Equal([]CardValue{ACE, JACK, SIX, FIVE, TWO}, hand.tiebreakerCardsValue)
}
