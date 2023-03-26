package arraysandhashing

func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, x := range nums{
        if val, ok := m[target - x]; ok{
            return []int{val, i}
        }else{
            m[x] = i
        }
    }

    return []int{0,0}
}
