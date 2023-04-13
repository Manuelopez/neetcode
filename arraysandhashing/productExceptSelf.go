package arraysandhashing

func productExceptSelf(nums []int) []int {
    var ans = make([]int, len(nums))

    calc := 1
    
    for i, x := range nums{
        ans[i] = calc
        calc *= x
    }

    calc = 1
    for i := len(nums) - 1; i >= 0; i--{
        ans[i] *= calc
        calc *= nums[i]
    }



    return ans
}
