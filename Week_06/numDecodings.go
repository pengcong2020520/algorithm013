// func numDecodings(s string) int {
//     //类似于爬楼梯 一次一格或者两格  但是注意两格的时候值要小于26
//     //后两位>26: dp[i] = dp[i-1]
//     //后两位<26: dp[i] = dp[i-1] + dp[i-2]
//     bytes := []byte(s)
//     length := len(bytes)
//     if s[0] == '0' {
//         return 0
//     }
//     if length == 1 {
//         return 1 
//     }
//     dp := make([]int, length+1)
//     dp[0] = 1
//     dp[1] = 1
//     for i := 2; i < length; i++ {
//         if s[i-1] == '0' {
//             if s[i-2] == '1' || s[i-2] == '2' {
//                 dp[i] = dp[i-1] + dp[i-2]
//             } else {
//                 return 0
//             }
//         } else if s[i-1] <= '6' && (s[i-2] == '1' || s[i-2] == '2') || (s[i-1] > '6' && s[i-2] == '1') {
//             dp[i] = dp[i-1] + dp[i-2]
//         } else {
//             dp[i] = dp[i-1]
//         }
           
//     }
//     return dp[length]
// }

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, tmp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		case s[i] == '0':
			cur = pre
		case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			tmp = cur
			cur = cur + pre
			pre = tmp
		default:
			pre = cur
		}
	}
	return cur
}
