package goutils

import (
"testing"
//"fmt"
)

// ****************************** TESTS ********************************************

func TestAbbreviate(t *testing.T) {

	// Test 1
	in := "abcdefg"
	out := "abc..."
	maxWidth := 6

	if x, _ := Abbreviate(in, maxWidth); x != out {
		t.Errorf("Abbreviate(%v, %v) = %v, want %v", in, maxWidth, x, out)
	}

	// Test 2
	out = "abcdefg"
	maxWidth = 7

	if x, _ := Abbreviate(in, maxWidth); x != out {
		t.Errorf("Abbreviate(%v, %v) = %v, want %v", in, maxWidth, x, out)
	}


	// Test 3
	out = "a..."
	maxWidth = 4

	if x, _ := Abbreviate(in, maxWidth); x != out {
		t.Errorf("Abbreviate(%v, %v) = %v, want %v", in, maxWidth, x, out)
	}
}




func TestAbbreviateFull(t *testing.T) {

	// Test 1
	in := "abcdefghijklmno"
	out := "abcdefg..."
	offset := -1
	maxWidth := 10

	if x, _ := AbbreviateFull(in, offset, maxWidth); x != out {
		t.Errorf("AbbreviateFull(%v, %v, %v) = %v, want %v", in, offset, maxWidth, x, out)
	}

	// Test 2
	out = "...fghi..."
	offset = 5
	maxWidth = 10

	if x, _ := AbbreviateFull(in, offset, maxWidth); x != out {
		t.Errorf("AbbreviateFull(%v, %v, %v) = %v, want %v", in, offset, maxWidth, x, out)
	}


	// Test 3
	out = "...ijklmno"
	offset = 12
	maxWidth = 10


	if x, _ := AbbreviateFull(in, offset, maxWidth); x != out {
		t.Errorf("AbbreviateFull(%v, %v, %v) = %v, want %v", in, offset, maxWidth, x, out)
	}
}





func TestIsBlank(t *testing.T) {

	// Test 1
	str := ""
	out := true

	if x := IsBlank(str); x != out {
		t.Errorf("IndexOf(%v) = %v, want %v", str, x, out)
	}

	// Test 2
	str = "   "
	out = true

	if x := IsBlank(str); x != out {
		t.Errorf("IndexOf(%v) = %v, want %v", str, x, out)
	}

	// Test 3
	str = " abc "
	out = false

	if x := IsBlank(str); x != out {
		t.Errorf("IndexOf(%v) = %v, want %v", str, x, out)
	}
} 



func TestDeleteWhiteSpace(t *testing.T) {

	// Test 1
	str := " a b c "
	out := "abc"

	if x := DeleteWhiteSpace(str); x != out {
		t.Errorf("IndexOf(%v) = %v, want %v", str, x, out)
	}

} 