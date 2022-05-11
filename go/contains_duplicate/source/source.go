package source

func ContainsDuplicate(nums []int) bool {
    hashMap := make(map[int]bool)
    
    for _, num := range nums {
        if _, ok := hashMap[num]; !ok {
            hashMap[num]=true
        }
    }
    
  
    return len(hashMap) != len(nums)
}