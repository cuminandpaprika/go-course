package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainCallsBubble(t *testing.T) {
	// Arrange
	var buf bytes.Buffer

	main()

	// Assert
	expected := strconv.Quote("")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected behavior for main")
		t.Errorf("expected: %s\nactual:%s", expected, actual)
	}
}

func TestBubbleSortsSingle(t *testing.T) {
	a := assert.New(t)

	input := []int{1}
	expected := []int{1}
	actual := bubble(input)

	a.Equal(expected, actual, "Bubble Sorts Single Input")
	// Check it's a copy
}

func TestBubbleSortReturnsSorted(t *testing.T) {
	a := assert.New(t)

	testCases := map[string]struct {
		input []int
		want  []int
	}{
		"single":     {input: []int{1}, want: []int{1}},
		"duplicates": {input: []int{2, 1, 4, 1}, want: []int{1, 1, 2, 4}},
		"negative":   {input: []int{-1, 1, -4, 0}, want: []int{-4, -1, 0, 1}},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		t.Run(name, func(t *testing.T) {
			result := bubble(test.input)
			a.Equal(result, test.want, "bubble failed to sort")
		})
	}
}
