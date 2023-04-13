package arraysandhashing

func ContainsDuplicate(nums []int) bool{
    m := make(map[int]bool)   

    for _, x := range nums{
        if _, ok := m[x]; ok{
            return true
        }else{
            m[x] = true
        }
    }

    return false
}
