package twopointer

func TwoSum(numbers []int, target int) []int {
    i := 0
    j := len(numbers) - 1

    for i < j{
        curr := numbers[i] + numbers[j]
        if curr == target{
            return []int{i + 1, j + 1}
        }else if curr > target{
            j--
        }else {
            i++
        }


    }

    return []int{0,0}
}
