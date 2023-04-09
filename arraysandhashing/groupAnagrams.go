package arraysandhashing


// my current own solution O(n^3) * too slow *
func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)

    for _, str := range strs{
        a := [26]int{} 
        for _, c := range str{
            a[c - 'a'] += 1
        }
        m[a] = append(m[a], str) 
    }

    ans := make([][]string, 0)

    for _, val := range m{
        ans = append(ans, val)
    }

    return ans
    
}    

