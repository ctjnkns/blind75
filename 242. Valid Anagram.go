func isAnagram(s string, t string) bool {
    //https://stackoverflow.com/a/22689818
    ss := strings.Split(s, "")
    sort.Strings(ss)
    fmt.Println(ss)
    tt := strings.Split(t, "")
    sort.Strings(tt)
    s = strings.Join(ss, "")
    t = strings.Join(tt, "")
    if s == t {
        return true
    }
    return false
}
