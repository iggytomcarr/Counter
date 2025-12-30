package main

import "testing"

func TestCountWorlds(t *testing.T) {
	input := ""
	wants := 0

	result := countWords([]byte(input))

	if result != wants {
		t.Log("Failed for input : ", input)
		t.Log("Wanted : ", wants)
		t.Log("Got : ", result)
		t.Fail()
	}

	input2 := "the little brown fox jumped"
	wants2 := 5

	result2 := countWords([]byte(input2))

	if result2 != wants2 {
		t.Log("Failed for input : ", input2)
		t.Log("Wanted : ", wants2)
		t.Log("Got : ", result2)
		t.Fail()
	}

	input3 := " "
	wants3 := 0

	result3 := countWords([]byte(input3))

	if result3 != wants3 {
		t.Log("Failed for input : ", input3)
		t.Log("Wanted : ", wants3)
		t.Log("Got : ", result3)
		t.Fail()
	}
}
