package main

import (
	"testing"

	test_assert "github.com/stretchr/testify/assert"
)

func TestIsBetterThan_True(t *testing.T) {
	assert := test_assert.New(t)

	card1 := Card{value: "10", suit: "H"}
	card2 := Card{value: "4", suit: "H"}

	assert.True(card1.IsBetterThan(card2))
}

func TestIsBetterThan_False(t *testing.T) {
	assert := test_assert.New(t)

	card1 := Card{value: "J", suit: "H"}
	card2 := Card{value: "K", suit: "H"}

	assert.False(card1.IsBetterThan(card2))
}
