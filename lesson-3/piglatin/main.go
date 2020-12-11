package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const consonants = "bcdfghjklmnpqrstvwxz"
const suffix = "ay"
const validationRegexp = "[^a-zA-Z .,!?-()]"
const punctuationMarks = " .,!?-()"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter string: ")

	for scanner.Scan() {
		input := strings.Trim(scanner.Text(), " ")

		if !validInput(input) {
			fmt.Println("Invalid input. Please, use latin letters and punctuation marks only.")
			break
		}

		fmt.Println("Your result: ", translate(input))
		break
	}

	if scanner.Err() != nil {
		fmt.Println("Oops... Some error occurred: ", scanner.Text())
	}
}

func validInput(s string) bool {
	match, _ := regexp.MatchString(validationRegexp, s)

	return !match && len(s) > 0
}

func translate(s string) string {
	var translated string
	var lexeme string

	for _, v := range strings.Split(s, "") {
		if strings.Contains(punctuationMarks, v) {
			translated += translateToPigLatin(lexeme)
			translated += v
			lexeme = ""
		} else {
			lexeme += v
		}
	}

	translated += translateToPigLatin(lexeme)

	return translated
}

func translateToPigLatin(s string) string {
	if len(s) == 0 {
		return s
	}
	stringSlice := strings.Split(s, "")


	for i, v := range stringSlice {
		if strings.Contains(consonants, v) || strings.Contains(strings.ToUpper(consonants), v) {
			continue
		} else {
			stringSlice = append(stringSlice, stringSlice[:i]...)
			stringSlice = stringSlice[i:]

			break
		}
	}

	return strings.Join(stringSlice, "") + suffix
}
