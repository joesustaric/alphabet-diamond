package main

import (
	"bytes"
	"testing"
)

func TestNewDiamondInfo(t *testing.T) {
	var happyTests = []struct {
		input    DiamondInfo
		expected DiamondInfo
	}{
		{NewDiamondInfo("A", 0), DiamondInfo{MiddleLetter: "A", IsA: true, MiddleWidth: 0}},
		{NewDiamondInfo("Z", 47), DiamondInfo{MiddleLetter: "Z", IsA: false, MiddleWidth: 47}},
	}

	for _, test := range happyTests {
		if areNotEqualx(test.input, test.expected) {
			t.Errorf("Expected %s \n Got %s \n", test.expected, test.input)
		}
	}
}

//Test to ensure the output display is correct
func TestDrawTheDumbDiamond(t *testing.T) {
	// 	input := NewLetters("A", "B", "C", "D", "E")
	// 	expectedZ := `
	// 		                       A
	// 	                        B B
	// 	                       C   C
	// 	                      D     D
	// 	                     E       E
	// 	                    F         F
	// 	                   G           G
	// 	                  H             H
	// 	                 I               I
	// 	                J                 J
	// 	               K                   K
	// 	              L                     L
	// 	             M                       M
	// 	            N                         N
	// 	           O                           O
	// 	          P                             P
	// 	         Q                               Q
	// 	        R                                 R
	// 	       S                                   S
	// 	      T                                     T
	// 	     U                                       U
	// 	    V                                         V
	// 	   W                                           W
	// 	  X                                             X
	// 	 Y                                               Y
	// 	Z                                                 Z
	// 	 Y                                               Y
	// 	  X                                             X
	// 	   W                                           W
	// 	    V                                         V
	// 	     U                                       U
	// 	      T                                     T
	// 	       S                                   S
	// 	        R                                 R
	// 	         Q                               Q
	// 	          P                             P
	// 	           O                           O
	// 	            N                         N
	// 	             M                       M
	// 	              L                     L
	// 	               K                   K
	// 	                J                 J
	// 	                 I               I
	// 	                  H             H
	// 	                   G           G
	// 	                    F         F
	// 	                     E       E
	// 	                      D     D
	// 	                       C   C
	// 	                        B B
	// 		                       A
	// `

	// buffer := new(bytes.Buffer)
	//
	// DrawTheDumbDiamond(buffer, input)
	//
	// if buffer.String() != expected {
	// 	t.Errorf("got \n %s, Expected \n %s", buffer.String(), expected)
	// }

	var happyTests = []struct {
		input    DiamondInfo
		expected string
	}{
		{DiamondInfo{MiddleLetter: "A", IsA: true, MiddleWidth: 0}, "A"},
	}

	for _, test := range happyTests {
		buffer := new(bytes.Buffer)
		DrawD(buffer, test.input)
		if buffer.String() != test.expected {
			t.Errorf("Expected %s \n Got %s \n", test.expected, buffer.String())
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
