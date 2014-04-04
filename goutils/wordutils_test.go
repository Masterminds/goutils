package wordutils

import (
"testing"
"fmt"

)

// Markdown for Example Documentation link: [![GoDoc](https://godoc.org/github.com/aokoli/GoWordUtils?status.png)](https://godoc.org/github.com/aokoli/GoWordUtils)

// EXAMPLES


func ExampleCapitalizeFully() {

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING"
	delimiters := []rune{' ', '.'} 

	fmt.Println (CapitalizeFully(in, delimiters...))
	// Output: 
	// Test Is Going Well.Thank.You.For Inquiring
}


func ExampleWrap_1() {

	in := "Bob Manuel Bob Manuel"
	wrapLength := 10

	fmt.Println (wordutils.Wrap(in, wrapLength))
	// Output: 
	// Bob Manuel
	// Bob Manuel
}


func ExampleWrap_2() {

	in := "BobManuelBob Bob"
	wrapLength := 10
	newLineStr := "<b\\>"
	wrapLongWords := false

	fmt.Println (wordutils.WrapCustom(in, wrapLength, newLineStr, wrapLongWords))
	// Output: 
	// BobManuelBob<b\\>Bob
}


func ExampleWrap_3() {

	in := "BobManuelBob Bob"
	wrapLength := 10
	newLineStr := "<b\\>"
	wrapLongWords := true

	fmt.Println (wordutils.WrapCustom(in, wrapLength, newLineStr, wrapLongWords))
	// Output: 
	// BobManuelB<b\\>ob Bob
}



// TESTS

func TestWrapNormalWord(t *testing.T) {

	in := "Bob Manuel Bob Manuel"
	out := "Bob Manuel\nBob Manuel"
	wrapLength := 10

	// const in, out = str1, str2 // TODO var vs const?
	if x := Wrap(in, wrapLength); x != out {
		t.Errorf("Wrap(%v) = %v, want %v", in, x, out)
	}
}


func TestWrapCustomLongWordFalse(t *testing.T) {

	in := "BobManuelBob Bob"
	out := "BobManuelBob<b\\>Bob"
	wrapLength := 10
	newLineStr := "<b\\>"
	wrapLongWords := false

	if x := WrapCustom(in, wrapLength, newLineStr, wrapLongWords); x != out {
		t.Errorf("Wrap(%v) = %v, want %v", in, x, out)
	}
}


func TestWrapCustomLongWordTrue(t *testing.T) {

	in := "BobManuelBob Bob"
	out := "BobManuelB<b\\>ob Bob"
	wrapLength := 10
	newLineStr := "<b\\>"
	wrapLongWords := true

	if x := WrapCustom(in, wrapLength, newLineStr, wrapLongWords); x != out {
		t.Errorf("WrapCustom(%v) = %v, want %v", in, x, out)
	}
}


func TestCapitalize(t *testing.T) {


	// Test 1: Checks if function works with 1 parameter, and default whitespace delimiter
	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING" 
	out := "TEsT IS GoiNG WELL.tHaNk.yOU.for InqUIrING" 

	if x := Capitalize(in); x != out {
		t.Errorf("Capitalize(%v) = %v, want %v", in, x, out)
	}


	// Test 2: Checks if function works with both parameters, with param 2 containing whitespace and '.'
	out = "TEsT IS GoiNG WELL.THaNk.YOU.For InqUIrING"	
	delimiters := []rune{' ', '.'} 

	if x := Capitalize(in, delimiters...); x != out {
		t.Errorf("Capitalize(%v) = %v, want %v", in, x, out)
	}
}



func TestCapitalizeFully(t *testing.T) {

	// Test 1
	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING" 
	out := "Test Is Going Well.thank.you.for Inquiring" 

	if x := CapitalizeFully(in); x != out {
		t.Errorf("CapitalizeFully(%v) = %v, want %v", in, x, out)
	}


	// Test 2
	out = "Test Is Going Well.Thank.You.For Inquiring"	
	delimiters := []rune{' ', '.'} 

	if x := CapitalizeFully(in, delimiters...); x != out {
		t.Errorf("CapitalizeFully(%v) = %v, want %v", in, x, out)
	}
}


func TestUncapitalize(t *testing.T) {

	// Test 1: Checks if function works with 1 parameter, and default whitespace delimiter
	in := "This Is A.Test" 
	out := "this is a.Test" 

	if x := Uncapitalize(in); x != out {
		t.Errorf("Uncapitalize(%v) = %v, want %v", in, x, out)
	}

	// Test 2: Checks if function works with both parameters, with param 2 containing whitespace and '.'
	out = "this is a.test" 
	delimiters := []rune{' ', '.'} 

	if x := Uncapitalize(in, delimiters...); x != out {
		t.Errorf("Uncapitalize(%v) = %v, want %v", in, x, out)
	}
}


func TestSwapCase(t *testing.T) {

	in := "This Is A.Test" 
	out := "tHIS iS a.tEST" 

	if x := SwapCase(in); x != out {
		t.Errorf("SwapCase(%v) = %v, want %v", in, x, out)
	}
}


func TestInitials(t *testing.T) {

	// Test 1
	in := "John Doe.Ray" 
	out := "JD" 

	if x := Initials(in); x != out {
		t.Errorf("Initials(%v) = %v, want %v", in, x, out)
	}


	// Test 2
	out = "JDR" 
	delimiters := []rune{' ','.'}

	if x := Initials(in, delimiters...); x != out {
		t.Errorf("Initials(%v) = %v, want %v", in, x, out)
	}
}