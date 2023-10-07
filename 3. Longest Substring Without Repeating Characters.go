func longestConsecutive(nums []int) int {
    numsHash := map[int]bool{} 
    longest := 0
    for _, num := range nums {
        numsHash[num] = true
    }
    for _, num := range nums {
        if _, ok := numsHash[num - 1] ; !ok {
            length := 0
            for numsHash[num + length] {
                length++
            }
            if length > longest { 
                longest = length
            }
        }
    }
    return longest
}
