/* DISREGARD THIS FILE. IT IS USED FOR RUNNING QUICK TESTS DURING PRODUCTION PHASE*/

package main

import (
	"fmt"
//	"strings" // http://golang.org/pkg/strings/
	// "bytes"  // http://golang.org/pkg/bytes/
//	"unicode" // http://golang.org/pkg/unicode/
//	"time"
//	"math/rand"
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

*/

func main() {

	chars := []rune {'1','2','9'}

	rand1, err1 := goutils.Random (-1, 0, 0, true, true)
	rand2, err2 := goutils.Random (10, 5, 1, true, true, chars...)

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

