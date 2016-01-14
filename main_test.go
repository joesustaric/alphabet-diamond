package main

import "testing"

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

func TestGetLettersForTheDiamond(t *testing.T) {

	var happyTests = []struct {
		input    string
		expected Letters
	}{
		{"E", Letters{First: "A", Second: "B", Third: "C", Fourth: "D", Fifth: "E"}},
		{"A", Letters{First: "W", Second: "X", Third: "Y", Fourth: "Z", Fifth: "A"}},
		{"B", Letters{First: "X", Second: "Y", Third: "Z", Fourth: "A", Fifth: "B"}},
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
