func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1
    sRune := []rune(s)
    for l < r {
        left := unicode.ToLower(sRune[l])
        right := unicode.ToLower(sRune[r])
        
        if !unicode.IsLetter(left) && !unicode.IsDigit(left) {
            l++
            continue
        }
        
        if !unicode.IsLetter(right) && !unicode.IsDigit(right) {
            r--
            continue
        }

        if left != right {
            return false
        }
        l++
        r--
    }
    return true 
}
