func threeSum(nums []int) [][]int {
    var result [][]int
    sort.Ints(nums)
    for i, num := range nums[:len(nums) - 2] {
        if i > 0 && num == nums[i - 1] {
            continue
        }
        l := i + 1
        r := len(nums) - 1
        for l < r {
            total := num + nums[l] + nums[r]
            if total == 0 {
				result = append(result, []int{num, nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l - 1] {
                    l++
                }
            } else if total > 0 {
                r--
            } else if total < 0 {
                l++
            }
        }
    }
    return result
}
