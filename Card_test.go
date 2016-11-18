package main

import (
	"testing"

	test_assert "github.com/stretchr/testify/assert"
)

func TestIsBetterThan_True(t *testing.T) {
	assert := test_assert.New(t)

	card1 := Card{Value: TEN, Suit: "H"}
	card2 := Card{Value: FOUR, Suit: "H"}

	assert.True(card1.IsBetterThan(card2))
}

func TestIsBetterThan_False(t *testing.T) {
	assert := test_assert.New(t)

	card1 := Card{Value: JACK, Suit: "H"}
	card2 := Card{Value: KING, Suit: "H"}

	assert.False(card1.IsBetterThan(card2))
}
