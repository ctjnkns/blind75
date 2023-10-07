func characterReplacement(s string, k int) int {
    charDict := make(map[byte]int)
    maxLen := 0
    l := 0;
    maxFreq := 0    
    for r := range s {
        charDict[s[r]] = charDict[s[r]] + 1
        if charDict[s[r]] > maxFreq {
            maxFreq = charDict[s[r]]
        }
        for (r - l + 1 - maxFreq) > k {
            charDict[s[l]] -=  1
            l++
        }
        if r - l + 1 > maxLen {
            maxLen = r - l + 1
        }
    }
    return maxLen
}
