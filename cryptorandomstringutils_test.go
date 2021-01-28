package goutils

import (
	"regexp"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestCryptoRandomNonAlphaNumeric(t *testing.T) {
	// If asked for a string 0 characters long, CryptoRandomNonAlphaNumeric should provide an empty string.
	if x, _ := CryptoRandomNonAlphaNumeric(0); utf8.RuneCountInString(x) != 0 {
		t.Errorf("String should be 0 characters; string was %v characters", utf8.RuneCountInString(x))
	}

	// Test CryptoRandomNonAlphaNumeric's ability to generate strings 1 through 100 characters in length.
	for i := 1; i < 101; i++ {
		if x, _ := CryptoRandomNonAlphaNumeric(i); utf8.RuneCountInString(x) != i {
			t.Errorf("String should be %v characters; string was %v characters", i, utf8.RuneCountInString(x))
		}
	}
}

func TestCryptoRandomAscii(t *testing.T) {
	// If asked for a string 0 characters long, CryptoRandomAscii should provide an empty string.
	if x, _ := CryptoRandomAscii(0); len(x) != 0 {
		t.Errorf("String should be 0 characters; string was %v characters", len(x))
	}

	// Test CryptoRandomAscii's ability to generate strings 1 through 100 characters in length.
	for i := 1; i < 101; i++ {
		if x, _ := CryptoRandomAscii(i); len(x) != i {
			t.Errorf("String should be %v characters; string was %v characters", i, len(x))
		}
	}
}

func TestCryptoRandomNumeric(t *testing.T) {
	// If asked for a string 0 characters long, CryptoRandomNumeric should provide an empty string.
	if x, _ := CryptoRandomNumeric(0); len(x) != 0 {
		t.Errorf("String should be 0 characters; string was %v characters", len(x))
	}

	// Test CryptoRandomNumeric's ability to generate strings 1 through 100 characters in length.
	for i := 1; i < 101; i++ {
		if x, _ := CryptoRandomNumeric(i); len(x) != i {
			t.Errorf("String should be %v characters; string was %v characters", i, len(x))
		}
	}
}

func TestCryptoRandomAlphabetic(t *testing.T) {
	// If asked for a string 0 characters long, CryptoRandomAlphabetic should provide an empty string.
	if x, _ := CryptoRandomAlphabetic(0); len(x) != 0 {
		t.Errorf("String should be 0 characters; string was %v characters", len(x))
	}

	// Test CryptoRandomAlphabetic's ability to generate strings 1 through 100 characters in length.
	for i := 1; i < 101; i++ {
		if x, _ := CryptoRandomAlphabetic(i); len(x) != i {
			t.Errorf("String should be %v characters; string was %v characters", i, len(x))
		}
	}
}

func TestCryptoRandomAlphaNumeric(t *testing.T) {
	// If asked for a string 0 characters long, CryptoRandomAlphaNumeric should provide an empty string.
	if x, _ := CryptoRandomAlphaNumeric(0); len(x) != 0 {
		t.Errorf("String should be 0 characters; string was %v characters", len(x))
	}

	// Test CryptoRandomAlphaNumeric's ability to generate strings 1 through 100 characters in length.
	for i := 1; i < 101; i++ {
		if x, _ := CryptoRandomAlphaNumeric(i); len(x) != i {
			t.Errorf("String should be %v characters; string was %v characters", i, len(x))
		}
	}
}

func TestCryptoRandAlphaNumeric_FuzzOnlyNumeric(t *testing.T) {

	// Testing for a reported regression in which some versions produced
	// a predictably small set of chars.
	iters := 1000
	charlen := 0
	for i := 0; i < 16; i++ {
		numOnly := 0
		charlen++
		for i := 0; i < iters; i++ {
			out, err := CryptoRandomAlphaNumeric(charlen)
			if err != nil {
				t.Fatal("func failed to produce a random thinger")
			}
			if _, err := strconv.Atoi(out); err == nil {
				numOnly++
			}

			m, err := regexp.MatchString("^[0-9a-zA-Z]+$", out)
			if err != nil {
				t.Fatal(err)
			}
			if !m {
				t.Fatal("Character is not alphanum")
			}
		}

		if numOnly == iters {
			t.Fatalf("Got %d numeric-only random sequences", numOnly)
		}
	}

}
