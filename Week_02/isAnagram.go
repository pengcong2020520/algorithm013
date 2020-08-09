func isAnagram(s string, t string) bool {
    
    if len(s) != len(t) {
        return false
    }
    length := len(s)
    strMap := make(map[rune]int, 26)
    for i := 0; i < length; i++ {
        strMap[rune(s[i])] += 1
        strMap[rune(t[i])] -= 1
    }
    for _, v := range strMap {
        if v != 0 {
            return false
        }
    }
    return true
}

