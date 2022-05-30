package source

func FindPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    
    for index, num := range nums {
        if index == 0 {
            if num > nums[index+1] {
                return index
            }
        } else if index == len(nums) - 1 {
            if num > nums[len(nums)-2] {
                return index
            }
        } else if num > nums[index + 1] && num > nums[index - 1] {
            return index
        }
    }
    
    return 0
}