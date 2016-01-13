package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

var letter *string

const (
	inputStringSize = 1
)

func init() {
	letter = flag.String("letter", "", "The letter which is the middle of the diamond.")
	flag.Parse()
}

func main() {
	fmt.Printf("letter = %s \n", *letter)
}

func Parse(input string) (string, error) {
	r := strings.TrimSpace(input)

	if notCorrectLength(r) {
		return "", fmt.Errorf("too many characters %s", r)
	} else if notAlphabetCharacter(r) {
		return "", fmt.Errorf("not a-z or A-Z character")
	}

	return strings.ToUpper(r), nil
}

func notCorrectLength(input string) bool {
	if len(input) != inputStringSize {
		return true
	}
	return false
}

func notAlphabetCharacter(input string) bool {
	matched, err := regexp.MatchString("[a-z|A-Z]", input)
	if !matched || err != nil {
		return true
	}
	return false
}
