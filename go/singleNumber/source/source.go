package source

func SingleNumber(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    
    setSlice := make(map[int]int)
    
    for _, num := range nums {
        setSlice[num] = setSlice[num] + 1
    }
    
    for _, num := range nums {
        if count := setSlice[num]; count == 1 {
            return num
        }
    }
    
    return -1
}