func countSubstrings(s string) int {
    //dp方程
    count := 0 
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    bytes := []byte(s)
    for right := 0; right < len(s); right++ {
        for left := 0; left <= right; left++ {
            if bytes[left] == bytes[right] && right - left <= 2{
                dp[left][right] = true
                count++
            } else if right - left > 2 && dp[left+1][right-1] && bytes[left] == bytes[right] {
                dp[left][right] = true
                count++
            }
        }
    }   
    return count
}