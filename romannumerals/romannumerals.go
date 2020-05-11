package romannumerals

func ToRoman(n int) string {
  if n == 1 { 
		return "I" 
	} else if n == 2 {
		return "II"
	} else if n == 3 {
		return "III"
	} else if n == 4 {
		return "IV"
	} else if n == 5 {
		return "V"
	}
	return "Unkown"
}