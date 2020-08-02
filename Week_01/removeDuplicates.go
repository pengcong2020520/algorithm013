func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    preNum := nums[0]
    count := 1
    for _, v := range nums {
        if v != preNum{
            preNum = v
            nums[count] = v 
            count += 1
        }
    }
    return count
}