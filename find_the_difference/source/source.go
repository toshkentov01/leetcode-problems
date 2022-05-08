package source

// This problem easily solved using ASCII
func FindTheDifference(s string, t string) byte {
    sumS := 0
    sumT := 0
    
    for _, data := range s {
        sumS += int(data)
    }
    
    for _, data := range t {
        sumT += int(data)
    }
    
    return byte(sumT - sumS)
}