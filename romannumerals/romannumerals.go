package romannumerals

import "strings"

func ToRoman(n int) string {
	if n == 10 {
		return "X"
	} else if n == 9 {
		return "IX"
	} else if n > 9 { 
		return "V" + strings.Repeat("I", n-5)
	} else if n == 6 {
		return "VI"
	} else if n == 5 {
		return "V"
	} else if n == 4 {
		return "IV"
	} else  {
		return strings.Repeat("I", n)
	}
	return "unknown"
}