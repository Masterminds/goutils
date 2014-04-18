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

package goutils


import (
	"fmt"
	"unicode" 
	"bytes"
	//"strings"
)



func Abbreviate (str string, maxWidth int) (string, error) {
	return AbbreviateFull(str, 0, maxWidth)
}

func AbbreviateFull (str string, offset int, maxWidth int) (string, error) {
    if str == "" {
        return "", nil
    }
    if maxWidth < 4 {
    	err := fmt.Errorf("stringutils illegal argument: Minimum abbreviation width is 4") 
    	return "", err
    }
    if len(str) <= maxWidth {
        return str, nil
    }
    if offset > len(str) {
        offset = len(str)
    }
    if len(str) - offset < (maxWidth - 3) {	// 15 - 5 < 10 - 3 =  10 < 7
        offset = len(str) - (maxWidth - 3)  
    }
    abrevMarker := "..."
    if offset <= 4 {
        return str[0:maxWidth - 3] + abrevMarker, nil// str.substring(0, maxWidth - 3) + abrevMarker;
    }
    if maxWidth < 7 {
    	err := fmt.Errorf("stringutils illegal argument: Minimum abbreviation width with offset is 7") 
    	return "", err
    }
    if (offset + maxWidth - 3) < len(str) {  // 5 + (10-3) < 15 = 12 < 15
    	abrevStr, _ := Abbreviate(str[offset:len(str)], (maxWidth - 3))
        return abrevMarker + abrevStr, nil// abrevMarker + abbreviate(str.substring(offset), maxWidth - 3);
    }
    return abrevMarker + str[(len(str) - (maxWidth - 3)):len(str)], nil // abrevMarker + str.substring(str.length() - (maxWidth - 3));
}




func DeleteWhiteSpace(str string) string {
        if str == "" {
            return str
        }
        sz := len(str)
        var chs bytes.Buffer 
        count := 0
        for i := 0; i < sz; i++ {
            ch := rune(str[i])
            if !unicode.IsSpace(ch) {
                chs.WriteRune(ch)
                count++
            }
        }
        if count == sz {
            return str
        }
        return chs.String()
}


/*
public static int indexOfDifference(final CharSequence cs1, final CharSequence cs2) {
    if (cs1 == cs2) {
        return INDEX_NOT_FOUND;
    }
    if (cs1 == null || cs2 == null) {
        return 0;
    }
    int i;
    for (i = 0; i < cs1.length() && i < cs2.length(); ++i) {
        if (cs1.charAt(i) != cs2.charAt(i)) {
            break;
        }
    }
    if (i < cs2.length() || i < cs1.length()) {
        return i;
    }
    return INDEX_NOT_FOUND;
}





 public static String deleteWhitespace(final String str) {
        if (isEmpty(str)) {
            return str;
        }
        final int sz = str.length();
        final char[] chs = new char[sz];
        int count = 0;
        for (int i = 0; i < sz; i++) {
            if (!Character.isWhitespace(str.charAt(i))) {
                chs[count++] = str.charAt(i);
            }
        }
        if (count == sz) {
            return str;
        }
        return new String(chs, 0, count);
    }
*/

/*

IsBlank checks if a string is whitespace or empty (""). Observe the following behavior:

    goutils.IsBlank("")        = true
    goutils.IsBlank(" ")       = true
    goutils.IsBlank("bob")     = false
    goutils.IsBlank("  bob  ") = false
 
Parameter:
    str - the string to check

Returns:
    true - if the string is whitespace or empty ("")
*/
func IsBlank(str string) bool {
   // return strings.TrimSpace(str) == ""
    strLen := len(str)
        if str == "" || strLen == 0 {
            return true
        }
        for i := 0; i < strLen; i++ {
            if unicode.IsSpace(rune(str[i])) == false {
                return false
            }
        }
    return true
}