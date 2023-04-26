package twopointer

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {return nums[i] < nums[j]})    
    ans := make([][]int, 0)
    for i, _ := range nums{
        if i > 0 && nums[i] == nums[i-1]{
                continue
        }
        l := i + 1
        r := len(nums) - 1
        for l < r{
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0{
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l++
                for nums[l] == nums[l-1] && l < r{
                    l++
                }
            }else if sum > 0{
                r--
            }else{
                l++
            }
        }
    }
    return ans
}
