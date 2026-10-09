package goutils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"testing"
	"unicode"
)

// ****************************** TESTS ********************************************

func TestRandomSeed(t *testing.T) {

	// count, start, end, letters, numbers := 5, 0, 0, true, true
	random := rand.New(rand.NewSource(10))
	out := "3ip9v"

	// Test 1: Simulating RandomAlphaNumeric(count int)
	if x, _ := RandomSeed(5, 0, 0, true, true, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, true, true, nil, random, x, out)
	}

	// Test 2: Simulating RandomAlphabetic(count int)
	out = "MBrbj"

	if x, _ := RandomSeed(5, 0, 0, true, false, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, true, false, nil, random, x, out)
	}

	// Test 3: Simulating RandomNumeric(count int)
	out = "88935"

	if x, _ := RandomSeed(5, 0, 0, false, true, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, false, true, nil, random, x, out)
	}

	// Test 4: Simulating RandomAscii(count int)
	out = "H_I;E"

	if x, _ := RandomSeed(5, 32, 127, false, false, nil, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 32, 127, false, false, nil, random, x, out)
	}

	// Test 5: Simulating RandomSeed(...) with custom chars
	chars := []rune{'1', '2', '3', 'a', 'b', 'c'}
	out = "2b2ca"

	if x, _ := RandomSeed(5, 0, 0, false, false, chars, random); x != out {
		t.Errorf("RandomSeed(%v, %v, %v, %v, %v, %v, %v) = %v, want %v", 5, 0, 0, false, false, chars, random, x, out)
	}

}

// ****************************** EXAMPLES ********************************************

func ExampleRandomSeed() {

	var seed int64 = 10 // If you change this seed #, the random sequence below will change
	random := rand.New(rand.NewSource(seed))
	chars := []rune{'1', '2', '3', 'a', 'b', 'c'}

	rand1, _ := RandomSeed(5, 0, 0, true, true, nil, random)      // RandomAlphaNumeric (Alphabets and numbers possible)
	rand2, _ := RandomSeed(5, 0, 0, true, false, nil, random)     // RandomAlphabetic (Only alphabets)
	rand3, _ := RandomSeed(5, 0, 0, false, true, nil, random)     // RandomNumeric (Only numbers)
	rand4, _ := RandomSeed(5, 32, 127, false, false, nil, random) // RandomAscii (Alphabets, numbers, and other ASCII chars)
	rand5, _ := RandomSeed(5, 0, 0, true, true, chars, random)    // RandomSeed with custom characters

	fmt.Println(rand1)
	fmt.Println(rand2)
	fmt.Println(rand3)
	fmt.Println(rand4)
	fmt.Println(rand5)
	// Output:
	// 3ip9v
	// MBrbj
	// 88935
	// H_I;E
	// 2b2ca
}

func TestRandAlphaNumeric_FuzzOnlyNumeric(t *testing.T) {

	// Testing for a reported regression in which some versions produced
	// a predictably small set of chars.
	iters := 1000
	charlen := 0
	for i := 0; i < 16; i++ {
		numOnly := 0
		charlen++
		for i := 0; i < iters; i++ {
			out, err := RandomAlphaNumeric(charlen)
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

func TestRandomNumeric(t *testing.T) {
	// If asked for a string 0 characters long, RandomNumeric should provide an empty string.
	if x, err := RandomNumeric(0); err != nil || len(x) != 0 {
		t.Errorf("RandomNumeric(0) = (%q, %v), want (\"\", nil)", x, err)
	}

	// Negative length should return an error.
	if _, err := RandomNumeric(-1); err == nil {
		t.Errorf("RandomNumeric(-1) expected error, got nil")
	}

	// Test RandomNumeric's ability to generate strings 1 through 100 characters in length.
	for i := 1; i <= 100; i++ {
		x, err := RandomNumeric(i)
		if err != nil {
			t.Fatalf("RandomNumeric(%d) unexpected error: %v", i, err)
		}
		if len(x) != i {
			t.Errorf("RandomNumeric(%d) length = %d, want %d", i, len(x), i)
		}
		if x[0] == '0' {
			t.Errorf("RandomNumeric(%d) generated leading zero: %s", i, x)
		}
		for _, ch := range x {
			if !unicode.IsDigit(ch) {
				t.Errorf("RandomNumeric(%d) contained non-digit character: %c", i, ch)
			}
		}
	}

	// Fuzz test multiple iterations across lengths to confirm no leading zeros and valid base-10 parsing.
	for _, length := range []int{1, 2, 3, 5, 10} {
		for i := 0; i < 1000; i++ {
			x, err := RandomNumeric(length)
			if err != nil {
				t.Fatalf("RandomNumeric(%d) failed at iteration %d: %v", length, i, err)
			}
			if len(x) != length {
				t.Fatalf("RandomNumeric(%d) length = %d, want %d", length, len(x), length)
			}
			if x[0] == '0' {
				t.Fatalf("RandomNumeric(%d) has leading zero: %s", length, x)
			}

			// Verify base-0 parsing succeeds and matches decimal parsing (avoids octal interpretation)
			val0, err0 := strconv.ParseInt(x, 0, 64)
			if err0 != nil {
				t.Fatalf("RandomNumeric(%d) %q failed base-0 ParseInt: %v", length, x, err0)
			}
			val10, err10 := strconv.ParseInt(x, 10, 64)
			if err10 != nil {
				t.Fatalf("RandomNumeric(%d) %q failed base-10 ParseInt: %v", length, x, err10)
			}
			if val0 != val10 {
				t.Fatalf("RandomNumeric(%d) %q parsed differently in base-0 (%d) vs base-10 (%d)", length, x, val0, val10)
			}
		}
	}
}

