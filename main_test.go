package main

import "testing"

func TestCountWorlds(t *testing.T) {
	input := "one two three four five"
	wants := 5

	result := countWords([]byte(input))

	if result != wants {
		t.Fail()
	}
}
