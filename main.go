package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func init() {
	letter = flag.String("letter", "", "The letter which is the middle of the diamond.")
	flag.Parse()
}

var letter *string
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	inputStringSize      = 1
	maxLettersForDiamond = 5
	diamondTemplate      = `
    {{.First}}
   {{.Second}} {{.Second}}
  {{.Third}}   {{.Third}}
 {{.Fourth}}     {{.Fourth}}
{{.Fifth}}       {{.Fifth}}
 {{.Fourth}}     {{.Fourth}}
  {{.Third}}   {{.Third}}
   {{.Second}} {{.Second}}
    {{.First}}

`
)

type Letters struct {
	First, Second, Third, Fourth, Fifth string
}

func NewLetters(first, second, third, fourth, fifth string) Letters {
	return Letters{First: first, Second: second, Third: third, Fourth: fourth, Fifth: fifth}
}

func main() {

	if input, err := Parse(*letter); err != nil {
		fmt.Println(err.Error())
	} else {
		l := GetDiamondLetters(input)

		DrawTheDumbDiamond(os.Stdout, l)
	}

}

func GetDiamondLetters(seed string) Letters {

	splitAlphabet := strings.Split(alphabet, seed)

	letters := []string{seed}

	letters = figureOutLettersForDiamond(splitAlphabet[0], letters)

	if len(letters) != maxLettersForDiamond {
		letters = figureOutLettersForDiamond(splitAlphabet[1], letters)
	}
	return NewLetters(letters[4], letters[3], letters[2], letters[1], letters[0])
}

func DrawTheDumbDiamond(out io.Writer, letters Letters) {
	t, err := template.New("diamond").Parse(diamondTemplate)
	if err != nil {
		panic(err)
	}
	t.Execute(out, letters)
}

func figureOutLettersForDiamond(letterSplitSet string, letters []string) []string {
	if len(letters) < maxLettersForDiamond {
		for i := len(letterSplitSet) - 1; i >= 0; i-- {
			letters = append(letters, string(letterSplitSet[i]))
			if len(letters) == 5 {
				return letters
			}
		}
	}
	return letters
}

func Parse(input string) (string, error) {
	r := strings.TrimSpace(input)

	if notCorrectLength(r) {
		return "", fmt.Errorf("input incorrect %s", r)
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
