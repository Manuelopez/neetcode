package twopointer

import (
    "regexp"
    "unicode"
    "strings"
)
func IsPalindrome(s string) bool {
   str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "") 
   str = strings.ToLower(str)
   str = SpaceMap(str)

   i := 0;
   j := len([]rune(str)) - 1

   for i < j{
       if str[i] != str[j]{
           return false
       }

       i++
       j--
   }

   return true


}

//helper function
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

