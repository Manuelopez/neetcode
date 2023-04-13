# Contains Duplicate

- Make a set to store values of array
- Loop through array storing the values "X" in the set
    - if value "X" already exist in the set return true
- Loop ends return false
- Time Complexity: O(n) 
- Space Complexity: O(n)
- [implementation](./containsdupplicate.go)

# Valid Anagram 

- Check if lenngths of string are equal
    - if not return false
- Create an array of with length of all the characters in the alphabet
    - This will be used to count the frequncy of each character 
    - Position 0 of the array is considered the character 'a' and positon 25 is considered the 'z' character
    - The array length is 26 of type int called "freq"
- Make a loop that goes from 0 to the length of the strings. Indexed by variable "i"
    - This would be the same cause we already checked
    - Grab the character at index "i" lets call "c"
    - Characters have value as per ascII standard so we can subtratct "c" - 'a' giving us a value from 0 to 25 
        - 'a' - 'a' == 0
        - 'z' - 'a' ==25
        - we do this for both string string s and t ging us indexex "x" and "y"
    - Add 1 to the "freq" array in position "x" and subtract 1 from the same array at position "y"
        - freq[x] += 1 
        = freq[y] -= 1
- Then loop through the "freq" array and try to find if all the values are 0, indexed by "i"
    - If value at postion "i" is not equal to 0 return false
- If loop ends without returning return true
- Time Complexity: O(n)
- Space Complexity: O(1)
- [implementation](./isAnagram.go) 


# Two Sum

- Make a HashMap with key as int and value as int 
    - The key of the map is going to be the value on the array
    - The value of the map is going to be the index where they key was at on the array
    - called "m"
- Loop through the array indexed by "i" and the values are "x"
    - subtract x from the target called "c"
    - check if c is found in the map
    - if value c is found in the map then return the value from that key and the current index "i" as an array
    - else (value "c" not found in map) add the value "x" to the map with the key of x and value of "i" its index
- If Loop ends without returning return a array with the values [0,0]
- Time Complexity: O(n)
- Space Complexity: O(n)
- [implementation](./twoSum.go)

# Group Anagrams

- Make a HashMap with key as [26]int and value as []string
- loop thorugh the array of strings
    - create a new array of 26 elemetns long of type int called "a"
    - loop through the string
        - create a feq array of the string using the "a" array
            - a[str[i] - 'a'] +=1
    - append the string to the map with key "a" (freq array just created)
- make a array of arrays of strings [][]string called "ans"
- loop through the map
    - append the arrays on the map to the new array "ans"
- return "ans"
- Time Complexty: O(n * m)
    - n: elements of array
    - m: average length of strings in array
- Space Complexity: O(n)
- [implementation](./groupAnagrams.go)

# Top K Frequent Elements

- Create a map "m" to count the frequency of the number key in value int 
- Create an array of array of ints called "c" with the parent array being of length of input + 1
- Loop through "input" giving the value "x"
    - add 1 to the map on key "x"
        - m[x] +=1
- loop through the map "m"
    - Append the key of the map on the "c" array at posiont "i" (the i position is the value from given key of the map)
        - append(c[val], key)
- create a "ans" array of interger
- loop backwards through the "c" array
    - append the values of the array to the "ans" array
        - append(ans, c[i]...)
    - if length of ans == k return ans
- if loop end without returning return ans
- Time Complexity:  O(n)
- Space Complexity: O(n)
- [implementation](./topkfrequentelements.go)

# Product of Array Except Self

- create an "ans" array length equal to the "input" array length
- create a "calc" variable of type int equal to 1
- loop through the "input" array indexed by "i" and giving value "x"
    - at positon "i" insert the value of "calc" int the "ans" array
    - cauculate "calc" to be equal to "calc" * "x"
        - calc *= x
- make calc equal to 1 again
- loop through array backwards index by "i" and giving value "x"
    - at positon "i" in the "ans" array make it equal to the product of itself time "calc"
        - ans[i] *= calc
    - cauculate "calc" to be equal to "calc" * "x"
        - calc *= x
- return ans
- Time Complexity:  O(n)
- Space Complexity: O(n)
- [implementation](./productExceptSelf.go)





