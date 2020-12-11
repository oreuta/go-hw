package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const validationRegexp = "[^a-zA-Z0-9 .,!?-()]"

var encodeVowels = map[rune]rune{
	'a': '1',
	'e': '2',
	'i': '3',
	'o': '4',
	'u': '5',
	'A': '1',
	'E': '2',
	'I': '3',
	'O': '4',
	'U': '5',
}

var decodeVowels = map[rune]rune{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u',
}

func main() {
	var mode string
	fmt.Print("Do you want to decode string or encode it? [encode/decode]: ")
	_, err := fmt.Scanln(&mode)

	if err != nil || (mode != "encode" && mode != "decode") {
		fmt.Print("There are two options only. Please, enter a valid one.")
		return
	}

	fmt.Print("Enter your string: ")
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for scanner.Scan() {
		input = scanner.Text()
		break
	}

	input = strings.Trim(input, " ")

	if err != nil || !validInput(input) {
		fmt.Print("Error. Please, make sure you use latin characters, numbers or punctuation characters only.")
		return
	}

	switch mode {
	case "encode":
		fmt.Println("Your result: ", swapString(input, encodeVowels))
	case "decode":
		fmt.Println("Your result: ", swapString(input, decodeVowels))
	}

}

func validInput(s string) bool {
	match, _ := regexp.MatchString(validationRegexp, s)

	return !match && len(s) > 0
}

func swapString(s string, mapping map[rune]rune) string {
	return strings.Map(func(r rune) rune {
		value, ok := mapping[r]

		if ok == false {
			return r
		}

		return value
	}, s)
}
