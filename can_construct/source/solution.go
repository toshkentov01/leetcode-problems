package source

func CanConstruct(ransomNote string, magazine string) bool {
    magazineMap := make(map[byte]int)
    
    for i := range magazine {
        magazineMap[magazine[i]] += 1
    }
    
    for i := range ransomNote {
        if value, ok := magazineMap[ransomNote[i]]; ok && value >= 1 {
            magazineMap[ransomNote[i]] -= 1
            continue
        }
    
        return false
    }
    
    return true
}