func containsDuplicate(nums []int) bool {
    numsHash := map[int]struct{}{} //struct is 0 bytes
    for _, num := range nums {
        if _, ok := numsHash[num]; ok { //ok is set to true if the key exists
            return true //return true if the value is already in the hash
        } else {
            numsHash[num] = struct{}{}
        }
    }
    return false
}
