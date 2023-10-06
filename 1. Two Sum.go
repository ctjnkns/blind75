func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if j, ok := numMap[(target - nums[i])] ; ok {
            return []int{i, j}
        }
        numMap[nums[i]] = i
    }
    return []int{0, 0}
}
