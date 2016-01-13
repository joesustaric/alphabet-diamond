package main

import "testing"

func TestParse(t *testing.T) {

	var happyTests = []struct {
		input    string // input
		expected string // expected result
	}{
		{"a", "A"},
		{"B", "B"},
		{"c", "C"},
		{"x", "X"},
		{"Y", "Y"},
		{"z", "Z"},
		{"g ", "G"},
		{" j", "J"},
		{" K ", "K"},
	}

	for _, test := range happyTests {
		result := Parse(test.input)
		if result != test.expected {
			t.Errorf("Fail Input=%s ,expected=%s ,result=%v \n",
				test.input, test.expected, result)
		}
	}

}
