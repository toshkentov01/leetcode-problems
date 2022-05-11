package source 

func MajorityElement(nums []int) int {
    hashMap := make(map[int]int)
     
    for _, num := range nums {
        hashMap[num] += 1
    }
    
    for key, value := range hashMap {
        if value > len(nums) / 2 {
            return key
        }
    }
                               
    return 0
}