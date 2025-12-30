package main

import "testing"

func TestCountWorlds(t *testing.T) {

	type testCase struct {
		name  string
		input string
		wants int
	}

	testCases := []testCase{
		{"zero", "", 0},
		{"five words", "the little brown fox jumped", 5},
		{"one space, but zero words", " ", 0},
		{"one word", "one", 1},
		{"new lines", "one two three\n four five", 5},
		{"spaces in between", "one  two   three four five", 5},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			result := countWords([]byte(tc.input))

			if result != tc.wants {
				t.Log("Name : ", tc.name)
				t.Log("Input : ", tc.input)
				t.Log("Wanted : ", tc.wants)
				t.Log("Got : ", result)
				t.Fail()
			}
		})

	}

}
