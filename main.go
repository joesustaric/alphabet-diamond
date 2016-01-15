package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

var letter *string

func init() {
	letter = flag.String("i", "", "The letter which is the middle of the diamond.")
	flag.Parse()
}

const (
	alphabet            = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	inputStringSize     = 1
	newLine             = "\n"
	letterA             = "A"
	letterAPaddingCalc  = -2
	widthOffset         = 2
	widthMultiplyFactor = 2
)

//DiamondInfo contains all the information to print the diamond. The middle letter
//The width of the middle of the diamond and 1 edge case which is A.
type DiamondInfo struct {
	MiddleWidth  int
	MiddleLetter string
	IsLetterA    bool
}

//NewDiamondInfo will return a new Diamond Info Object which is used for displaying
//the diamond. It will autocalculate the width nessecary for the midpoint.
//Also will flag it as an A as a special exception.
func NewDiamondInfo(middleLetter string) DiamondInfo {
	if middleLetter == letterA {
		return DiamondInfo{MiddleLetter: middleLetter,
			MiddleWidth: 0,
			IsLetterA:   true}
	}
	return DiamondInfo{MiddleLetter: middleLetter,
		MiddleWidth: figureOutTheWidth(middleLetter),
		IsLetterA:   false}
}

func figureOutTheWidth(letter string) int {
	letterNumber := strings.Index(alphabet, letter)
	return (letterNumber * widthMultiplyFactor) - widthOffset
}

func main() {
	start := time.Now()

	if input, err := Parse(*letter); err != nil {
		fmt.Println(err.Error())
	} else {
		DrawJamesAPrettyDiamond(os.Stdout, NewDiamondInfo(input))
	}

	duration := time.Now().Sub(start)
	fmt.Printf("Joes app internal method timing = %s \n", duration)
}

//DrawJamesAPrettyDiamond will draw a diamond with the alphabet to the given
//writer. Must give it a Diamond info to seed it.
func DrawJamesAPrettyDiamond(out io.Writer, di DiamondInfo) {
	if di.IsLetterA {
		out.Write([]byte(letterA + newLine))
		return
	}

	splitAlphabet, lPadding, dBottom := strings.Split(alphabet, di.MiddleLetter), 0, ""
	middleLine := newLine + printDiamondLine(lPadding, di.MiddleWidth, di.MiddleLetter)

	for i := len(splitAlphabet[0]) - 1; i >= 0; i-- {
		lPadding++
		l := string(splitAlphabet[0][i])
		dBottom = dBottom + printDiamondLine(lPadding, figureOutTheWidth(l), l)
	}

	dTop := newLine + reverse(dBottom)

	out.Write([]byte(dTop + middleLine + dBottom + newLine))
}

//reverse any string
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[j], r[i] = r[i], r[j]
	}
	return string(r)
}

func printDiamondLine(leftPaddingSize, widthPaddingSize int, letter string) string {
	boundaryPadding, widthPadding := generatePadding(leftPaddingSize, widthPaddingSize)

	if isPaddingForLetterA(widthPaddingSize) {
		return boundaryPadding + letter + boundaryPadding
	}

	return boundaryPadding + letter + widthPadding + letter + boundaryPadding + newLine
}

func generatePadding(leftPaddingSize, widthPaddingSize int) (string, string) {
	boundryPadding, widthPadding := "", ""

	for i := 0; i < leftPaddingSize; i++ {
		boundryPadding = boundryPadding + " "
	}

	for i := 0; i <= widthPaddingSize; i++ {
		widthPadding = widthPadding + " "
	}
	return boundryPadding, widthPadding
}

func isPaddingForLetterA(paddingSize int) bool {
	return paddingSize == letterAPaddingCalc
}

//Parse takes a string input and will return the sanitised string alphabet
//character. It will return an error if the input was not parsed correctly.
func Parse(input string) (string, error) {
	r := strings.TrimSpace(input)

	if notCorrectLength(r) || notAlphabetCharacter(r) {
		return "", fmt.Errorf("INVALID INPUT")
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
