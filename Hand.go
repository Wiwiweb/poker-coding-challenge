package main

import (
	"fmt"
	"sort"
)

type Hand struct {
	Cards                [5]Card
	combo                HandCombo
	tiebreakerCardsValue []CardValue
}

var handComboOrder []HandCombo

func init() {
	flush := HandCombo{value: FLUSH, isThisCombo: isFlush, getTiebreakers: getFlushTiebreakers}
	threeOfAKind := HandCombo{value: THREE_OF_A_KIND, isThisCombo: isThreeOfAKind, getTiebreakers: getThreeOfAKindTiebreakers}
	pair := HandCombo{value: PAIR, isThisCombo: isPair, getTiebreakers: getPairTiebreakers}
	highCard := HandCombo{value: HIGH_CARD, isThisCombo: isHighCard, getTiebreakers: getHighCardTiebreakers}

	handComboOrder = []HandCombo{flush, threeOfAKind, pair, highCard}
}

func (h *Hand) calculateHandValue() {
	for _, combo := range handComboOrder {
		if combo.isThisCombo(h.Cards) {
			h.combo = combo
			h.tiebreakerCardsValue = combo.getTiebreakers(h.Cards)
			break
		}
	}
}

func isFlush(cards [5]Card) bool {
	flushSuite := cards[0].Suit
	for _, card := range cards {
		if card.Suit != flushSuite {
			return false
		}
	}
	return true
}

func getFlushTiebreakers(cards [5]Card) []CardValue {
	cardValues := getCardValues(cards)
	reverseSort(cardValues)
	return cardValues
}

func isThreeOfAKind(cards [5]Card) bool {
	count := make(map[CardValue]int)
	for _, card := range cards {
		count[card.Value]++
		if count[card.Value] == 3 {
			return true
		}
	}
	return false
}

// Three of a kinds can't tie, beyond the value of the three of a kind
// (because that would require 6 cards of the same value)
// So they only need the one tiebreaker value
func getThreeOfAKindTiebreakers(cards [5]Card) []CardValue {
	count := make(map[CardValue]int)
	for _, card := range cards {
		count[card.Value]++
		if count[card.Value] == 3 {
			return []CardValue{card.Value}
		}
	}
	// If we reach that point, then there wasn't a three of a kind in the first place
	// This shouldn't happen!
	return []CardValue{}
}

// Slightly different from three of a kind, because we can have two pairs
// but we only want the highest one
func isPair(cards [5]Card) bool {
	count := make(map[CardValue]int)
	var highestPairValue CardValue
	for _, card := range cards {
		count[card.Value]++
		if count[card.Value] == 2 {
			highestPairValue = maxCardValue(highestPairValue, card.Value)
		}
	}

	if highestPairValue != 0 {
		return true
	}
	return false
}

func getPairTiebreakers(cards [5]Card) []CardValue {
	count := make(map[CardValue]int)
	var highestPairValue CardValue
	for _, card := range cards {
		count[card.Value]++
		if count[card.Value] == 2 {
			highestPairValue = maxCardValue(highestPairValue, card.Value)
		}
	}

	fmt.Println("highest card:", highestPairValue)
	tiebreakers := []CardValue{highestPairValue}
	cardValues := getCardValuesExcept(cards, highestPairValue)
	reverseSort(cardValues)
	tiebreakers = append(tiebreakers, cardValues...)
	return tiebreakers
}

// If we reach that point, all other combos have been ruled out, so High Card is automatically true
func isHighCard(_ [5]Card) bool {
	return true
}

func getHighCardTiebreakers(cards [5]Card) []CardValue {
	cardValues := getCardValues(cards)
	reverseSort(cardValues)
	return cardValues
}

func getCardValues(cards [5]Card) []CardValue {
	cardValues := []CardValue{}
	for _, card := range cards {
		cardValues = append(cardValues, card.Value)
	}
	return cardValues
}

func getCardValuesExcept(cards [5]Card, value CardValue) []CardValue {
	cardValues := []CardValue{}
	for _, card := range cards {
		if card.Value != value {
			cardValues = append(cardValues, card.Value)
		}
	}
	return cardValues
}

func reverseSort(s CardsByValue) {
	sort.Sort(sort.Reverse(s))
}

func maxCardValue(a, b CardValue) CardValue {
	if a < b {
		return b
	}
	return a
}
