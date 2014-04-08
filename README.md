GoUtils
===========

GoUtils provides users with utility functions to manipulate strings in various ways. It is a Go implementation of the 
string manipulation libraries of Java Apache Commons. GoUtils includes the following Java Apache Commons libraries:
* WordUtils	
* RandomStringUtils  

## Installation
If you have Go set up on your system, enter this within the command line/terminal:

	go get github.com/aokoli/goutils
    
If you do not have Go set up on your system, please follow the [Go installation directions from the documenation](http://golang.org/doc/install), and then follow the instructions above to install GoUtils.


## Documentation 
GoUtils doc is available here: [![GoDoc](https://godoc.org/github.com/aokoli/GoUtils/goutils?status.png)](https://godoc.org/github.com/aokoli/GoUtils/goutils)


## Usage
The code snippet below shows an example on how to use a library within GoUtils. In this instance, and the function is the `Initials` function located within the `wordutils.go` library file.

    package main
    
    import (
        "fmt"
    	"github.com/aokoli/goutils"
    )
    
    func main() {
    	// EXAMPLE 1: A goutils function which returns no errors
        fmt.Println (goutils.Initials("John Doe Foo")) // Prints out "JDF"

        // EXAMPLE 2: A goutils function which returns an error
        // Intialize two variables, one for the expected string, and one for the potential error
        rand1, err1 := goutils.Random (-1, 0, 0, true, true)  

        if err1 != nil { 
			fmt.Println(err1) // Prints out error message because -1 was entered as the first parameter in goutils.Random(...)
		} else {
			fmt.Println(rand1) 
		}

    }


## License
GoUtils is licensed under the Apache License, Version 2.0. Please check the LICENSE.txt file or visit http://www.apache.org/licenses/LICENSE-2.0 for a copy of the license. 

## Website
* [GoUtils webpage](#)

## Mailing List
Contact [okolialex@gmail.com](mailto:okolialex@mail.com) to be added to the mailing list.

