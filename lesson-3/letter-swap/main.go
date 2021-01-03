package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var validationRegexp = regexp.MustCompile("[^a-zA-Z0-9 .,!?-()]") // this is a heavy operation. let's do it only once

type vowelMapping map[rune]rune // we use it more than one time, so let's define it as a custom type

var encodeVowels = vowelMapping{
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

var decodeVowels = vowelMapping{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u',
}

var mapping = map[string]vowelMapping{ // it's a static data, so using a predefined mapping we can avoid switch (additional complexity in run-rime)
	"encode" : encodeVowels,
	"decode" : decodeVowels,
}

func main() {
	var mode string
	fmt.Print("Do you want to decode string or encode it? [encode/decode]: ")
	_, err := fmt.Scanln(&mode)

	if err != nil || (mode != "encode" && mode != "decode") { // it'd be better to separate technical and logical issues (err and data checks)
		fmt.Print("There are two options only. Please, enter a valid one.")
		return
	}

	fmt.Print("Enter your string: ")
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for scanner.Scan() {
		input = scanner.Text()
		break // if you want to read only one line, just don't use for-loop
	}

	input = strings.Trim(input, " ")

	if !validInput(input) {
		fmt.Print("Error. Please, make sure you use latin characters, numbers or punctuation characters only.")
		return
	}

	fmt.Println("Your result: ", swapString(input, mapping[mode]))
}

func validInput(s string) bool {
	if s == "" { // to avoid making a heavy regexp operation if it is not needed
		return false
	}
 	return validationRegexp.MatchString(s)
}

func swapString(s string, mapping vowelMapping) string {
	return strings.Map(func(r rune) rune {
		value, ok := mapping[r]

		if !ok {
			return r
		}

		return value
	}, s)
}
