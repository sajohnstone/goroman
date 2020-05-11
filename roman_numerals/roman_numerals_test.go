package roman_numerals

import "testing"

func TestToRoman(t *testing.T){
  expected := "I"
  actual := ToRoman(1)
  if actual != expected {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}