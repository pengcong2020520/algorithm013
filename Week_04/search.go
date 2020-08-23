func search(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    mid := 0
    for left <= right {
        mid = (left+right)/2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[left] {  
            //nums[mid]<target<nums[left]  在前半部分找
            if nums[mid] > target && target >= nums[left]{  //前半部分
                right = mid - 1
            } else {    //后半部分
                left = mid + 1
            }
        } else {                  //在前半部分  （找出在后半部分找的情况）
            // nums[right]>target>nums[mid]   在后半部分找
            if nums[mid] < target && target <= nums[right]{
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
