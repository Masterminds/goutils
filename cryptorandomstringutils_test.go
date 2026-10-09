package goutils

import (
	"regexp"
	"strconv"
	"testing"
	"unicode"
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
	if x, err := CryptoRandomNumeric(0); err != nil || len(x) != 0 {
		t.Errorf("CryptoRandomNumeric(0) = (%q, %v), want (\"\", nil)", x, err)
	}

	// Negative length should return an error.
	if _, err := CryptoRandomNumeric(-1); err == nil {
		t.Errorf("CryptoRandomNumeric(-1) expected error, got nil")
	}

	// Test CryptoRandomNumeric's ability to generate strings 1 through 100 characters in length.
	for i := 1; i <= 100; i++ {
		x, err := CryptoRandomNumeric(i)
		if err != nil {
			t.Fatalf("CryptoRandomNumeric(%d) unexpected error: %v", i, err)
		}
		if len(x) != i {
			t.Errorf("CryptoRandomNumeric(%d) length = %d, want %d", i, len(x), i)
		}
		if x[0] == '0' {
			t.Errorf("CryptoRandomNumeric(%d) generated leading zero: %s", i, x)
		}
		for _, ch := range x {
			if !unicode.IsDigit(ch) {
				t.Errorf("CryptoRandomNumeric(%d) contained non-digit character: %c", i, ch)
			}
		}
	}

	// Fuzz test multiple iterations across lengths to confirm no leading zeros and valid base-10 parsing.
	for _, length := range []int{1, 2, 3, 5, 10} {
		for i := 0; i < 1000; i++ {
			x, err := CryptoRandomNumeric(length)
			if err != nil {
				t.Fatalf("CryptoRandomNumeric(%d) failed at iteration %d: %v", length, i, err)
			}
			if len(x) != length {
				t.Fatalf("CryptoRandomNumeric(%d) length = %d, want %d", length, len(x), length)
			}
			if x[0] == '0' {
				t.Fatalf("CryptoRandomNumeric(%d) has leading zero: %s", length, x)
			}

			// Verify base-0 parsing succeeds and matches decimal parsing (avoids octal interpretation)
			val0, err0 := strconv.ParseInt(x, 0, 64)
			if err0 != nil {
				t.Fatalf("CryptoRandomNumeric(%d) %q failed base-0 ParseInt: %v", length, x, err0)
			}
			val10, err10 := strconv.ParseInt(x, 10, 64)
			if err10 != nil {
				t.Fatalf("CryptoRandomNumeric(%d) %q failed base-10 ParseInt: %v", length, x, err10)
			}
			if val0 != val10 {
				t.Fatalf("CryptoRandomNumeric(%d) %q parsed differently in base-0 (%d) vs base-10 (%d)", length, x, val0, val10)
			}
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
