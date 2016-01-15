package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

var letter *string

func init() {
	letter = flag.String("i", "", "The letter which is the middle of the diamond.")
	flag.Parse()
}

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

//Letters is used as an input for the template for the diamond output.
type Letters struct {
	First, Second, Third, Fourth, Fifth string
}

//NewLetters will give you a new Letter object.
func NewLetters(first, second, third, fourth, fifth string) Letters {
	return Letters{First: first,
		Second: second,
		Third:  third,
		Fourth: fourth,
		Fifth:  fifth}
}

//DiamondInfo contains all the information to print the diamond. The middle letter
//The width of the middle of the diamond and 1 edge case which is A.
type DiamondInfo struct {
	MiddleWidth  int
	MiddleLetter string
	IsA          bool
}

func NewDiamondInfo(middleLetter string, middleWidth int) DiamondInfo {
	if middleLetter == "A" {
		return DiamondInfo{MiddleLetter: middleLetter, MiddleWidth: middleWidth, IsA: true}
	}
	return DiamondInfo{MiddleLetter: middleLetter, MiddleWidth: middleWidth, IsA: false}
}

func main() {
	start := time.Now()

	if input, err := Parse(*letter); err != nil {
		fmt.Println(err.Error())
	} else {
		DrawTheDumbDiamond(os.Stdout, GetDiamondLetters(input))
	}

	duration := time.Now().Sub(start)
	fmt.Printf("Done in %s. Boom!\n", duration)
}

func DrawD(out io.Writer, diamondInfo DiamondInfo) {
	if diamondInfo.IsA {
		out.Write([]byte("A"))
		return
	}

	splitAlphabet := strings.Split(alphabet, diamondInfo.MiddleLetter)
	leftPadding := 0
	middle := "\n" + printDiamondLine(leftPadding, diamondInfo.MiddleWidth, diamondInfo.MiddleLetter)
	bottom := ""

	for i := len(splitAlphabet[0]) - 1; i >= 0; i-- {
		leftPadding++
		bottom = bottom + printDiamondLine(leftPadding, ((i*2)-2), string(splitAlphabet[0][i]))
	}

	top := "\n" + reverse(bottom)

	out.Write([]byte(top + middle + bottom + "\n"))
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[j], r[i] = r[i], r[j]
	}
	return string(r)
}

func printDiamondLine(leftPadding, widthPadding int, letter string) string {
	result := ""
	for i := 0; i < leftPadding; i++ {
		result = result + " "
	}
	rightPadding := result
	result = result + letter
	for i := 0; i <= widthPadding; i++ {
		result = result + " "
	}

	if widthPadding == -2 {
		return result + rightPadding
	}

	return result + letter + rightPadding + "\n"
}

//GetDiamondLetters will figure out from a input letter [A-Z] the other letters
//that need to be used to display in the diamond template.
//It returns a Letter obejct initalised with the correct values.
func GetDiamondLetters(inputLetter string) Letters {

	splitAlphabet, letters := strings.Split(alphabet, inputLetter), []string{inputLetter}

	for _, characterSet := range splitAlphabet {
		letters = figureOutLettersForDiamond(characterSet, letters)
	}

	if len(letters) == maxLettersForDiamond {
		return NewLetters(letters[4], letters[3], letters[2], letters[1], letters[0])
	}

	return Letters{}
}

//DrawTheDumbDiamond will draw James amazing diamond figure given it has an object
//which satisfies the writer interface. stdout / file / socket where ever you want
//to see this majestic diamond to fulfill all your diamond fantasy's.
func DrawTheDumbDiamond(out io.Writer, letters Letters) {

	if t, err := template.New("diamond").Parse(diamondTemplate); err != nil {
		panic(err)
	} else {
		t.Execute(out, letters)
	}

}

func figureOutLettersForDiamond(letterSplitSet string, letters []string) []string {
	//traverse backwards
	for i := len(letterSplitSet) - 1; i >= 0; i-- {
		if len(letters) == maxLettersForDiamond {
			return letters
		}
		letters = append(letters, string(letterSplitSet[i]))
	}
	return letters
}

//Parse takes a string input and will return the sanitised string alphabet
//character. It will return an error if the input was not parsed correctly.
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
	if matched, err := regexp.MatchString("[a-z|A-Z]", input); !matched || err != nil {
		return true
	}
	return false
}
