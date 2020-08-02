func twoSum(nums []int, target int) []int {
// //1. 循环查找解决
//时间复杂度O（n2）
//空间复杂度O（1）
// res := []int{}
// for i, _ := range nums {
//     for j, _ := range nums {
//         if nums[i] + nums[j] == target {
//             if i == j {
//                 continue
//             }
//             res = append(res, i)
//             res = append(res, j)
//             return res
//         }
//     }
// }
//     return res

//2. 定义一个map循环查找
//时间复杂度O（n）
//空间复杂度O（n）
    length := len(nums)
    m := make(map[int]int, length)
    res := []int{}
    for i, v := range nums {
        a, ok := m[target-v]
        if ok {
            res = append(res, a)
            res = append(res, i)
            return res
        }
        m[v] = i 
    }
    return res
}