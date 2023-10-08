func findMin(nums []int) int {
    res := nums[0]
    l := 0
    r := len(nums) - 1
    for l <= r {
        if nums[l] < nums[r] {
            if nums[l] < res {
                res = nums[l]
            }
        }
        m := (l + r) / 2
        if nums[m] < res {
            res = nums[m]
        }
        if nums[m] >= nums[l] {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return res
}
