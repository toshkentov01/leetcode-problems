package source

func CountWords(words1 []string, words2 []string) int {
    hashMap := make(map[string]int) 
    hashMap2 := make(map[string]int) 
    result := 0
    
    for _, word := range words1 {
        hashMap[word] += 1
    }
    
    for _, word := range words2 {
        hashMap2[word] += 1
    }
    
    for _, word := range words2 {
        if count, ok := hashMap[word]; ok && count == 1 {
            if count2, _ := hashMap2[word]; count2 == 1 {
                result += 1
            }
        }
    }
    
    return result
}