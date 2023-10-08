func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        m := (l + r) / 2
        if target == nums[m] {
            return m
        }
        if nums[l] <= nums[m] {
            //we're in the left sorted portion
            if target > nums[m] || target < nums[l] {
                l = m + 1
            } else { //target is less than m and greater than l
                r = m - 1
            }
        } else {
            //we're in the right sorted portion
            if target < nums[m] || target > nums[r]{
                r = m - 1
            } else { //target greater than m and less than r
                l = m + 1
            }
        }
    }
    return -1
}
