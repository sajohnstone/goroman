package romannumerals

import "testing"

func TestToRoman(t *testing.T){
	var testCases = map[string]int{
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
	}

	for expected, input := range testCases {
		actual := ToRoman(input)
		if actual != expected {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
		}
	}
}