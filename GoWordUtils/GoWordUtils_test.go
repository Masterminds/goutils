package gowordutils

import "testing"

// go test github.com/aokoli/gowordutils

func TestWrap(t *testing.T) {
	const in, out = "test string", "test string" // TODO var vs const?
	if x := wrap(in); x != out {
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
