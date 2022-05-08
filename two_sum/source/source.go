package source

// TwoSum ...
func TwoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    
    for index, num := range nums {
        numsMap[num] = index
    }
    for index, num := range nums {
        possibleAnswer := target - num
        
        if i, ok := numsMap[possibleAnswer]; i != index && ok {
            return []int{i, index}
        }
    }
    
    return []int{}
}