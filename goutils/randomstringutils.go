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

// Package goutils provides utility functions to manipulate strings in various ways.
package goutils


import (
	/*
	"bytes"   
	"strings"
	
	*/
	"fmt"
	//"errors"
	"unicode" 
	"math"
	"math/rand"
	"time"

)

var _ = rand.Intn(1)

func init() {
	 rand.Seed(time.Now().UnixNano()) // Provides the time-based seed used to generate random #s 

}

/*,  Random random*/
// /*
func Random (count int, start int, end int, letters bool, numbers bool, chars ...rune) (string, error) {

    if count == 0 {
        return "", nil
    } else if count < 0 {
    	//errMsg := "Requested random string length " + string(rune(count)) + " is less than 0."
    	// err := errors.New(errMsg)
    	err := fmt.Errorf("randomstringutils illegal argument: Requested random string length %v is less than 0.", count) // equiv to err := errors.New("...")
    	return "", err
    }
    if chars != nil && len(chars) == 0 {
    	err := fmt.Errorf("randomstringutils illegal argument: The chars array must not be empty") 
    	return "", err
    }

    if start == 0 && end == 0 {
        if chars != nil {
            end = len(chars)
        } else {
            if !letters && !numbers {
                end = math.MaxInt32
            } else {
                end = 'z' + 1;
                start = ' ';                
            }
        }
    } else {
        if end <= start {
	        err := fmt.Errorf("randomstringutils illegal argument: Parameter end (%v) must be greater than start (%v)", end, start) 
	    	return "", err
        }
    }

    buffer := make([]rune, count)
    gap := end - start

    for count != 0 {
    	count--
        var ch rune
        if chars == nil {
            ch = rune(rand.Intn(gap) + start)// TODO rune converstion needed? (char) (random.nextInt(gap) + start);
        } else {
            ch = chars[rand.Intn(gap) + start]// chars[random.nextInt(gap) + start];
        }
        if letters && unicode.IsLetter(ch) || numbers && unicode.IsDigit(ch) || !letters && !numbers { // TODO Check IsNumber
            if ch >= 56320 && ch <= 57343 { // TODO
                if count == 0 {
                   count++
                } else {
                    // low surrogate, insert high surrogate after putting it in
                    buffer[count] = ch
                    count--
                    buffer[count] = rune(55296 + rand.Intn(128))// (char) (55296 + random.nextInt(128));
                }
            } else if ch >= 55296 && ch <= 56191 { // TODO
                if count == 0 {
                   count++
                } else {
                    // high surrogate, insert low surrogate before putting it in
                    buffer[count] = rune(56320 + rand.Intn(128)) // (char) (56320 + random.nextInt(128));
                    count--
                    buffer[count] = ch
                }
            } else if ch >= 56192 && ch <= 56319 {
                // private high surrogate, no effing clue, so skip it // TODO
                count++
            } else {
                buffer[count] = ch
            }
        } else {
            count++
        }
    }
    return  string(buffer), nil
}
// */