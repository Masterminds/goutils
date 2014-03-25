/*
Copyright 2014 Alexander Okoli

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package gowordutils is a Golang implementation of the WordUtils class of
// Apache Commons, located within the package org.apache.commons.lang3.text
package gowordutils

import (
	// "fmt"
	"strings" // http://golang.org/pkg/strings/
	"bytes"  // http://golang.org/pkg/bytes/
	// "unicode" // http://golang.org/pkg/unicode/
)

///*
func Nilly(s string) string {

	 if s == "" {
            return "It was nil"
        }

	return "was not nil"
}
// */


// /*
func Wrap(str string, wrapLength int, newLineStr string, wrapLongWords bool) string {

	/*
		If columns/wrapLength = 10
		Start counting the words
		Add words to a line. If a word is above the 10 column count,
		send it to a new line 

	 */

	 if (str == "") {
            return ""
        }
        if (newLineStr == "") {
            newLineStr = "\n" // TODO Assumes "\n" is seperator. Explore SystemUtils.LINE_SEPARATOR;
        }
        if (wrapLength < 1) {
            wrapLength = 1;
        }
          

        inputLineLength := len(str) //final int inputLineLength = str.length();
        offset := 0 // int offset = 0;

        var wrappedLine bytes.Buffer  // final StringBuilder wrappedLine = new StringBuilder(inputLineLength + 32);
        
        // 1st loop: 10 - 0 > 9 true
        // 2nd loop: 10 - 5 > 9 false
        // 3rd loop: Say, 20 - 5 > 9 true

        for inputLineLength - offset > wrapLength { //   while (inputLineLength - offset > wrapLength) {
        	// If str = "Bob Manuel", the ff if statement is skipped
          	if rune(str[offset]) == ' ' {	// if (str.charAt(offset) == ' ') {
                offset++
                continue
            }


            // spaceToWrapAt = 4, as that is the index of the gap between "Bob" and "Manuel"
            // i.e. "Look for the next time ' ' occured after Bob"
            end := wrapLength + offset + 1
            spaceToWrapAt := strings.LastIndex(str[offset:end], " ") + offset // int spaceToWrapAt = str.lastIndexOf(' ', wrapLength + offset);


            // 4 > 0, so we go into if statement. (Why no spaceToWrapAt != -1? Since ' ' last index search would ALWAYS yield -1 for long words?)
            if spaceToWrapAt >= offset {  //if (spaceToWrapAt >= offset) {
            	// The following grabs "Bob", and appends "Bob" to buildLine
                // normal case
                wrappedLine.WriteString(str[offset:spaceToWrapAt]) // wrappedLine.append(str.substring(offset, spaceToWrapAt));
                // Appends the divider to buildLine, yielding "Bob\n"
                wrappedLine.WriteString(newLineStr) // wrappedLine.append(newLineStr);
                // offset = 4 + 1 = 5
                offset = spaceToWrapAt + 1;
      //      }

         
      //   } 

            
            /*
                
            } else {
                // really long word or URL
                if (wrapLongWords) {
                    // wrap really long word one line at a time
                    wrappedLine.append(str.substring(offset, wrapLength + offset));
                    wrappedLine.append(newLineStr);
                    offset += wrapLength;
                } else {
                    // do not wrap really long word, just extend beyond limit
                    spaceToWrapAt = str.indexOf(' ', wrapLength + offset);
                    if (spaceToWrapAt >= 0) {
                        wrappedLine.append(str.substring(offset, spaceToWrapAt));
                        wrappedLine.append(newLineStr);
                        offset = spaceToWrapAt + 1;
                    } else {
                        wrappedLine.append(str.substring(offset));
                        offset = inputLineLength;
                    }
                }
            */

            }
        }

        // Grabs "Manuel", and appends to "Bob\n", yielding Bob\nManuel
        // Whatever is left in line is short enough to just pass through
        wrappedLine.WriteString(str[offset:len(str)]) // wrappedLine.append(str.substring(offset));

        return wrappedLine.String() //return wrappedLine.toString();

}


// */








/*
// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) (float64, string) {
	z := 1.0                    // test comment 1
	for i := 0; i < 1000; i++ { // test comment 2
		z -= (z*z - x) / (2 * z)
	}
	a := "okay"
	return z, a
}

*/

// cd %gopath%/src/github.com/aokoli/main