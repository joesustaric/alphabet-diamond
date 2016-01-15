package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestNewDiamondInfo(t *testing.T) {
	var happyTests = []struct {
		input    DiamondInfo
		expected DiamondInfo
	}{
		{NewDiamondInfo("A", 0), DiamondInfo{MiddleLetter: "A", IsA: true, MiddleWidth: 0}},
		{NewDiamondInfo("Z", 48), DiamondInfo{MiddleLetter: "Z", IsA: false, MiddleWidth: 48}},
	}

	for _, test := range happyTests {
		if areNotEqualx(test.input, test.expected) {
			t.Errorf("Expected %s \n Got %s \n", test.expected, test.input)
		}
	}
}

//Test to ensure the output display is correct
func TestDrawTheDumbDiamond(t *testing.T) {
	var expectedE, expectedZ []byte
	var err error

	if expectedE, err = ioutil.ReadFile("testEDiamond.txt"); err != nil {
		t.Errorf("Could not read from file '%s'", "testEDiamond.txt")
	} else if expectedZ, err = ioutil.ReadFile("testZDiamond.txt"); err != nil {
		t.Errorf("Could not read from file '%s'", "testZDiamond.txt")
	}

	var happyTests = []struct {
		input    DiamondInfo
		expected string
	}{
		{DiamondInfo{MiddleLetter: "A", IsA: true, MiddleWidth: 0}, "A"},
		{DiamondInfo{MiddleLetter: "E", IsA: false, MiddleWidth: 6}, string(expectedE)},
		{DiamondInfo{MiddleLetter: "Z", IsA: false, MiddleWidth: 48}, string(expectedZ)},
	}

	for _, test := range happyTests {
		buffer := new(bytes.Buffer)
		DrawD(buffer, test.input)
		if test.expected != buffer.String() {
			t.Errorf("Expected \n(%s)\n Got \n(%s)\n", test.expected, buffer.String())
		}
	}

}

//Test to ensure the input argument is parsed correctly
func TestParse(t *testing.T) {

	var happyTests = []struct {
		input    string
		expected string
	}{
		{"a", "A"},
		{"B", "B"},
		{"c", "C"},
		{"x", "X"},
		{"Y", "Y"},
		{"z", "Z"},
		{"g ", "G"},
		{" j   ", "J"},
		{" K", "K"},
	}

	for _, test := range happyTests {
		result, err := Parse(test.input)
		if result != test.expected {
			t.Errorf("Fail Input=%s ,expected=%s ,result=%v \n",
				test.input, test.expected, result)
		}

		if err != nil {
			t.Errorf("Fail Input=%s ,expected no error but got one\n",
				test.input)
		}
	}

	var sadTests = []struct {
		input string
	}{
		{"aa"},
		{""},
		{"1"},
		{"123"},
		{"1a"},
		{"z1"},
		{"?"},
		{"z@"},
	}

	for _, test := range sadTests {
		_, err := Parse(test.input)
		if err == nil {
			t.Errorf("Fail Input=%s ,expected an error \n",
				test.input)
		}
	}
}

//Tests to figure out the letters to print on the diamond
func TestGetLettersForTheDiamond(t *testing.T) {
	var happyTests = []struct {
		input    string
		expected Letters
	}{
		{"E", Letters{First: "A", Second: "B", Third: "C", Fourth: "D", Fifth: "E"}},
		{"A", Letters{First: "W", Second: "X", Third: "Y", Fourth: "Z", Fifth: "A"}},
		{"B", Letters{First: "X", Second: "Y", Third: "Z", Fourth: "A", Fifth: "B"}},
		{"D", Letters{First: "Z", Second: "A", Third: "B", Fourth: "C", Fifth: "D"}},
		{"Z", Letters{First: "V", Second: "W", Third: "X", Fourth: "Y", Fifth: "Z"}},
	}

	for _, test := range happyTests {
		result := GetDiamondLetters(test.input)
		if areNotEqual(result, test.expected) {
			t.Errorf("Fail Input=%s ,expected=%s ,result=%v \n",
				test.input, test.expected, result)
		}
	}
}

func areNotEqualx(item1, item2 DiamondInfo) bool {
	if item1.IsA != item2.IsA {
		return true
	} else if item1.MiddleLetter != item2.MiddleLetter {
		return true
	} else if item1.MiddleWidth != item2.MiddleWidth {
		return true
	}
	return false
}

//a custom assertion for the Letter type
func areNotEqual(result Letters, expected Letters) bool {
	if result.First != expected.First {
		return true
	} else if result.Second != expected.Second {
		return true
	} else if result.Third != expected.Third {
		return true
	} else if result.Fourth != expected.Fourth {
		return true
	} else if result.Fifth != expected.Fifth {
		return true
	}

	return false
}
