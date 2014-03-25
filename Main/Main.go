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

	str := "Bob Manuel is"

	fmt.Println (strings.LastIndex(str[0:9+0], " "))
	fmt.Println (strings.LastIndex(str[9:12], " "))
	fmt.Println (strings.LastIndex(str[0:len(str)], " "))

	// Wrap display
	fmt.Println(gowordutils.Wrap("Bob Manuel Bob Manuel Bob Manuel", 10, "", false))
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