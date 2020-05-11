package main

import (
	"fmt"
	"strconv"
	"./romannumerals"
)

var appName = "The Romain Numerals Converter"

func main() {
	var inputValue string
	var inputValueAsInt int

	fmt.Printf("Welcome to %v\n", appName)
	fmt.Printf("Enter a numeric value: ")
	fmt.Scanln(&inputValue)
	inputValueAsInt, _ = strconv.Atoi(inputValue)
  fmt.Printf("> %v\n", roman_numerals.ToRoman(inputValueAsInt))
}