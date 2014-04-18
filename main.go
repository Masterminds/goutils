/* DISREGARD THIS FILE. IT IS USED FOR RUNNING QUICK TESTS DURING PRODUCTION PHASE*/

package main

import (
	"fmt"
//	"strings" // http://golang.org/pkg/strings/
	// "bytes"  // http://golang.org/pkg/bytes/
//	"unicode" // http://golang.org/pkg/unicode/
//	"time"
	"math/rand"
//	"math"
	"errors"
	"github.com/aokoli/goutils"
)

// var _ = strings.ToUpper("debug")
//var _ = unicode.IsPrint('4')
var _ = errors.New("nada")
var _ = goutils.Initials("John Doe Foo")

/*
Explain: Uncapitalize uncapitalizes all the DELIMETER separated words in a string. If no delimeter parameter is 
specified, the default would be a whitespace character. Whitespace is determined by unicode.IsSpace()



## Future Direction
The current aim of the project is to provide users with utility functions to manipulate strings in various ways; however, it might be expanded to include
more Java Apache Commons libraries. 

*/

func main() {

	// ******************** STRINGUTILS *******************
	ab, _ := goutils.Abbreviate("abcdefg", 6)
	ab2, _ := goutils.Abbreviate("", 6)
	fmt.Println(ab)
	fmt.Println(ab2)
	fmt.Println(goutils.IndexOf("abcdaf","a",0))

	/*

	StringUtils.abbreviate(null, *)      = null
 StringUtils.abbreviate("", 4)        = ""
 StringUtils.abbreviate("abcdefg", 6) = "abc..."
 StringUtils.abbreviate("abcdefg", 7) = "abcdefg"
 StringUtils.abbreviate("abcdefg", 8) = "abcdefg"
 StringUtils.abbreviate("abcdefg", 4) = "a..."
 StringUtils.abbreviate("abcdefg", 3) = IllegalArgumentException


	*/


	// ******************** RANDOMSTRINGUTILS *******************

	chars := []rune {'1','2','9'}

	//fmt.Println(math.MaxInt32)

	// random := rand.Seed(time.Now().UnixNano()) 
	// /*
	r := rand.New(rand.NewSource(2))
	r2 := rand.New(rand.NewSource(2))
	rand3, err3 := goutils.RandomSeed (5, 0, 0, true, true, chars, r)
	rand4, err4 := goutils.RandomSeed (5, 0, 0, true, true, chars, r2)

	rand5, err5 := goutils.Random (5, 0, 0, true, true, chars...)
	rand6, err6 := goutils.Random (5, 0, 0, true, true, chars...)

	rand7, err7 := goutils.Random (5, 0, 0, true, true)
	rand8, err8 := goutils.RandomAlphaNumericCustom (20, false, true)

	rand9, err9 := goutils.RandomAlphaNumeric (5)

	rand10, err10 := goutils.RandomAlphabetic (5)

	rand11, err11 := goutils.RandomAscii (5)
	rand12, err12 := goutils.RandomNonAlphaNumeric (5)
	rand13, err13 := goutils.RandomNumeric (10)



	 if err3 != nil {
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
		fmt.Println(err6)
		fmt.Println(err7)
		fmt.Println(err8)
		fmt.Println(err9)
		fmt.Println(err10)
		fmt.Println(err11)
		fmt.Println(err12)
		fmt.Println(err13)
	} else {
		fmt.Println(rand3)
		fmt.Println(rand4)
		fmt.Println(rand5)
		fmt.Println(rand6)
		fmt.Println(rand7)
		fmt.Println(rand8)
		fmt.Println(rand9)
		fmt.Println(rand10)
		fmt.Println(rand11)
		fmt.Println(rand12)
		fmt.Println(rand13)
	 }


	// */

	/*
	
	
	//time.Sleep(10)
	rand5, err5  := goutils.Random (5, 0, 0, true, true, chars...)
	rand6, err6  := goutils.Random (5, 0, 0, true, true, chars...)
	rand7, err7  := goutils.Random (5, 0, 0, true, true, chars...)

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(rand3)
	}

	if err4 != nil {
		fmt.Println(err4)
		fmt.Println(err5)
		fmt.Println(err6)
		fmt.Println(err7)
	} else {
		fmt.Println(rand4)
		fmt.Println(rand5)
		fmt.Println(rand6)
		fmt.Println(rand7)
	}
	*/
	
	 /*
	rand1, err1 := goutils.Random (5, 0, 0, true, true)
	rand2, err2 := goutils.Random (5, 94, 97, true, true, chars...)



	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(rand1)
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(rand2)
	}
	// */

	/* rand3, err3 := goutils.Random (10, 55296, 57343, true, true)  // Surr operations HANG.
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(rand3)
	}
	*/





	/*
	rand.Seed(time.Now().UnixNano()) // Try changing this number!
	answers := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])	
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])	
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	*/

	//str := "Bob Manuel is"
	/*r := make([]rune, 2)

	 r[0] = ' '
	 r[1] = '.'
	 */

	//delimiters := []rune{' ', 'a','b'} 


	//fmt.Println (goutils.Initials("John Doe Foo"))



/*

	// To convert variable v to type T, you would do: T(v)
	rn := []rune("foo")  // equiv "foo".toCharArray()
	fmt.Println(len(rn))
	str := string(rn)	// equiv String(new char[]) 
	fmt.Println(str)


	var buffer bytes.Buffer

    for i := 0; i < 10; i++ {
        buffer.WriteString("a")
    }

    fmt.Println(buffer.String())

   // char := rune("hello"[1]) // equivalent to String.CharAt  (string[index] converts ASCII, 117. The ASCII is then converted to rune)
    fmt.Println("hello"[1:4])


*/



	// a, b := wordutils.Sqrt(5)
	// fmt.Printf("Here is a: %v\n", a)
	// fmt.Printf("Here is b: %v\n", b)
    // fmt.Printf("Test print = %v, %v\n", "somto", 1) 
    // fmt.Println("This " + "That" + "\u00E4")   
}

