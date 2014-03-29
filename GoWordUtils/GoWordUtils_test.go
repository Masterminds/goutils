package gowordutils

import "testing"

// go test github.com/aokoli/gowordutils


func ExampleCapitalizeFullyCustom_nas() {

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING"
	out := "Test Is Going Well.Thank.You.For Inquiring"	
	delimiters := []rune{' ', '.'} 

	fmt.Println (CapitalizeFullyCustom(in, delimiters))
	// Output: 
	// Test Is Going Well.Thank.You.For Inquiring
}

func ExampleCapitalizeFully() {

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING"
	out := "Test Is Going Well.Thank.You.For Inquiring"	
	delimiters := []rune{' ', '.'} 

	fmt.Println (CapitalizeFullyCustom(in, delimiters))
	// Output: 
	// Test Is Going Well.Thank.You.For Inquiring
}


func Okay() {

}



/*
func TestWrapNormalWord(t *testing.T) {

	in := "Bob Manuel Bob Manuel"
	out := "Bob Manuel\nBob Manuel"
	wrapLength := 10

	// const in, out = str1, str2 // TODO var vs const?
	if x := Wrap(in, wrapLength); x != out {
		t.Errorf("Wrap(%v) = %v, want %v", in, x, out)
	}
}


func TestWrapLongWordFalse(t *testing.T) {

	in := "BobManuelBob Bob"
	out := "BobManuelBob\nBob"
	wrapLength := 10

	if x := Wrap(in, wrapLength); x != out {
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

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING" //"the test.is.going.well.thank you.for inquiring"
	out := "TEsT IS GoiNG WELL.tHaNk.yOU.for InqUIrING" // "The Test.is.going.well.thank You.for Inquiring"

	if x := Capitalize(in); x != out {
		t.Errorf("Capitalize(%v) = %v, want %v", in, x, out)
	}
}


func TestCapitalizeCustom(t *testing.T) {

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING"// "the test.is.going.well.thank you.for inquiring"
	out := "TEsT IS GoiNG WELL.THaNk.YOU.For InqUIrING"	// "The Test.Is.Going.Well.Thank You.For Inquiring"
	delimiters := []rune{' ', '.'} 

	if x := CapitalizeCustom(in, delimiters); x != out {
		t.Errorf("Capitalize(%v) = %v, want %v", in, x, out)
	}
}



func TestCapitalizeFully(t *testing.T) {

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING" 
	out := "Test Is Going Well.thank.you.for Inquiring" 

	if x := CapitalizeFully(in); x != out {
		t.Errorf("Capitalize(%v) = %v, want %v", in, x, out)
	}
}


func TestCapitalizeFullyCustom(t *testing.T) {

	in := "tEsT iS goiNG wELL.tHaNk.yOU.for inqUIrING"
	out := "Test Is Going Well.Thank.You.For Inquiring"	
	delimiters := []rune{' ', '.'} 

	if x := CapitalizeFullyCustom(in, delimiters); x != out {
		t.Errorf("Capitalize(%v) = %v, want %v", in, x, out)
	}
}
*/

