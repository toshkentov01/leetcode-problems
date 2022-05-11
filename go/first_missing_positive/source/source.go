package source

func FirstMissingPositive(nums []int) int {
    setStyleSlice := make(map[int]bool)
    mostMatchedAnswer := 1
    
    for _, num := range nums {
        if num > 0 {
            setStyleSlice[num] = true
        }
    }
    
    for {	
        if exists := setStyleSlice[mostMatchedAnswer]; !exists {
            return mostMatchedAnswer
        }
        
        mostMatchedAnswer += 1
    }
}