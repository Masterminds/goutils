package main

import (
	"fmt"
	"strings" // http://golang.org/pkg/strings/
	// "bytes"  // http://golang.org/pkg/bytes/
	"unicode" // http://golang.org/pkg/unicode/

	"github.com/aokoli/gowordutils"
)

var _ = strings.ToUpper("debug")
var _ = unicode.IsPrint('4')
// var _ = gowordutils.Nilly("debug")

func main() {

	//str := "Bob Manuel is"
	r := make([]rune, 2)

	r[0] = ' '
	r[1] = '.'

	delimiters := []rune{' ', 'a','b'} 


	fmt.Println (gowordutils.Capitalize("test is going well.thank you.for Inquiring", r))
	//fmt.Println (strings.IndexRune(str[9:12], ' '))
	//fmt.Println (strings.LastIndex(str[9:12], " "))
	//fmt.Println (strings.LastIndex(str[0:len(str)], " "))

	// Wrap display
	//fmt.Println(gowordutils.Wrap("Bob Manuel is my test text. I am trying to see how this thing works properly. If it works, that will be very great", 25))
	//fmt.Println(gowordutils.Wrap("Bob Manuel is here and there", 9, "<b\\>", false))

	







	
	//r := 0
	//fmt.Println(r == nil)














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



	// a, b := gowordutils.Sqrt(5)
	// fmt.Printf("Here is a: %v\n", a)
	// fmt.Printf("Here is b: %v\n", b)
    // fmt.Printf("Test print = %v, %v\n", "somto", 1) 
    // fmt.Println("This " + "That" + "\u00E4")   
}