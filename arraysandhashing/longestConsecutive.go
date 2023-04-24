package arraysandhashing

func LongestConsecutive(nums []int) int {
    m := make(map[int]bool)
    for _,x := range nums{
        m[x] = true
    }

    
    longest := 0
    length := 0
    for _, x := range nums{
        if _, ok := m[x-1]; !ok{
            for m[x+length]{
                length++
            }
            if length > longest{
                longest = length
            }
            length = 0
        }
    }

    return longest
}    
