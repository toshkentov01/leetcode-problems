package source

func SearchInsert(nums []int, target int) int {
    var result int
    
    if target < nums[0] {
        return 0
    
    } else if target > nums[len(nums) - 1] {
        return len(nums)
    }
    
    for index, num := range nums {
        if num == target {
            return index
        }
        
        if index != len(nums) - 1 && nums[index] < target && nums[index+1] > target {
            result = index + 1
        }
    }
    
    return result
}
