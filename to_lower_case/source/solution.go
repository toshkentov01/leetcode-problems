package source

func toLowerCase(s string) string {
    var result string
    
    for i := range s {
        if s[i] > 64 && s[i] < 91 {
            result += string(s[i] + 32)
            continue
        }
        
        result += string(s[i])
    }
    
    return result
}