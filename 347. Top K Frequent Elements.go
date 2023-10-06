func topKFrequent(nums []int, k int) (res []int) {
    count := map[int]int{}
    freq := make([][]int, len(nums) + 1)

    for _, n := range nums {
        if c, ok := count[n]; ok {
            count[n] = c + 1
        } else {
            count[n] = 1
        }
    }

    for n, c := range count {
        freq[c] = append(freq[c], n)
    }

    for i := len(freq) - 1; i > 0; i-- {
        res = append(res, freq[i]...)
        if len(res) == k {
            return res
        }
    }
    return nil
}
