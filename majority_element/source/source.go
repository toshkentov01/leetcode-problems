package source 

func MajorityElement(nums []int) int {
    nMap := make(map[int]int)

    for i := range nums {
        nMap[nums[i]]++
    }

    for key, value := range nMap {
        if value > (len(nums)/2) {
            return key
        }
    }

    return 0
}