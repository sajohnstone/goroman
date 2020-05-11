package romannumerals

import (
	"strconv"
	"testing"
)

func TestToRoman(t *testing.T){
	var testCases = map[string]int{
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
		"VI": 6,
		"IX": 9,
		"X": 10,
	}

	for expected, input := range testCases {
		actual := ToRoman(input)
		if actual != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s' input: '%s'", expected, actual, strconv.Itoa(input))
		}
	}
}