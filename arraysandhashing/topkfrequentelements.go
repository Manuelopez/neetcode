package arraysandhashing

import ("sort")


func TopKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    countSlice := make([][]int, len(nums) + 1)
    for _, x := range nums{
        m[x] += 1
    } 

    for num, rep := range m{
        countSlice[rep] = append(countSlice[rep], num)
    }

    a := make([]int, 0)

    for i:= len(countSlice)-1; i > 0; i--{
        if(len(a) == k){
            return a
        }

        a = append(a, countSlice[i]...)
    }

    return a
}


func TopKFrequentMyVersion(nums []int, k int) []int {
    m := make(map[int]int)
    for _, x := range nums{
        m[x] += 1
    } 

    a := make([]struct{val int
     rep int}, 0)

    for key, val := range m{
        a = append(a, struct{val int 
        rep int}{val: key, rep: val})
    }

    sort.Slice(a, func(i, j int)bool{
        return a[i].rep > a[j].rep
    })

    b := make([]int, 0)
    for i := 0; i < k; i++{
        b = append(b, a[i].val)
    }
    return b
}
