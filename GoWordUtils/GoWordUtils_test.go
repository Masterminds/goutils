package gowordutils

import "testing"

// go test github.com/aokoli/gowordutils

func TestWrapNormalWord(t *testing.T) {

	in := "Bob Manuel Bob Manuel"
	out := "Bob Manuel\nBob Manuel"
	wrapLength := 10
	newLineStr := ""
	wrapLongWords := false

	// const in, out = str1, str2 // TODO var vs const?
	if x := Wrap(in, wrapLength, newLineStr, wrapLongWords); x != out {
		t.Errorf("wrap(%v) = %v, want %v", in, x, out)
	}
}


func TestWrapLongWordFalse(t *testing.T) {

	in := "BobManuelBob Bob"
	out := "BobManuelBob\nBob"
	wrapLength := 10
	newLineStr := ""
	wrapLongWords := false

	if x := Wrap(in, wrapLength, newLineStr, wrapLongWords); x != out {
		t.Errorf("wrap(%v) = %v, want %v", in, x, out)
	}
}


func TestWrapLongWordTrue(t *testing.T) {

	in := "BobManuelBob Bob"
	out := "BobManuelB\nob Bob"
	wrapLength := 10
	newLineStr := ""
	wrapLongWords := true

	if x := Wrap(in, wrapLength, newLineStr, wrapLongWords); x != out {
		t.Errorf("wrap(%v) = %v, want %v", in, x, out)
	}
}

/*
func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
*/
