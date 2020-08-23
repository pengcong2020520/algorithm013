//1. 暴力法  遍历数组，记录每次可到达的最远距离
// func canJump(nums []int) bool {
//     var maxJump int
//     for i := 0; i < len(nums); i++ {
//         if i > maxJump {
//             return false
//         }
//         maxJump = max(maxJump, i+nums[i]) 
//     }
//     return true
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// 2. 贪心法
func canJump(nums []int) bool {
    length := len(nums)
    end := length - 1
    for i := length-1; i >= 0; i-- {
        if nums[i] + i >= end {
            end = i
        }
    }
    return end == 0
}