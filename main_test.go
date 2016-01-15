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
		{NewDiamondInfo("A"), DiamondInfo{MiddleLetter: "A", IsLetterA: true, MiddleWidth: 0}},
		{NewDiamondInfo("Z"), DiamondInfo{MiddleLetter: "Z", IsLetterA: false, MiddleWidth: 48}},
	}

	for _, test := range happyTests {
		if areNotEqual(test.input, test.expected) {
			t.Errorf("Expected %v \n Got %v \n", test.expected, test.input)
		}
	}
}

//Test to ensure the output display is correct
func TestDrawJamesAPrettyDiamond(t *testing.T) {
	var expectedE, expectedZ []byte
	var err error

	// load the expected output from files
	if expectedE, err = ioutil.ReadFile("testEDiamond.txt"); err != nil {
		t.Errorf("Could not read from file '%s'", "testEDiamond.txt")
	} else if expectedZ, err = ioutil.ReadFile("testZDiamond.txt"); err != nil {
		t.Errorf("Could not read from file '%s'", "testZDiamond.txt")
	}

	var happyTests = []struct {
		input    DiamondInfo
		expected string
	}{
		{DiamondInfo{MiddleLetter: "A", IsLetterA: true, MiddleWidth: 0}, "A\n"},
		{DiamondInfo{MiddleLetter: "E", IsLetterA: false, MiddleWidth: 6}, string(expectedE)},
		{DiamondInfo{MiddleLetter: "Z", IsLetterA: false, MiddleWidth: 48}, string(expectedZ)},
	}

	for _, test := range happyTests {
		buffer := new(bytes.Buffer)
		DrawJamesAPrettyDiamond(buffer, test.input)
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

//a custom assertion for the Letter type
func areNotEqual(item1, item2 DiamondInfo) bool {
	if item1.IsLetterA != item2.IsLetterA {
		return true
	} else if item1.MiddleLetter != item2.MiddleLetter {
		return true
	} else if item1.MiddleWidth != item2.MiddleWidth {
		return true
	}
	return false
}
