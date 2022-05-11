package source

func LengthOfLastWord(s string) int {
    lengthCounter := 0
    
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == 32 && lengthCounter == 0 {
            continue
        
        } else if s[i] == 32 {
            return lengthCounter
        } 
		        
        lengthCounter += 1
    }
    
    return lengthCounter
}