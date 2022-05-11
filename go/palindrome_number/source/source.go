package source

func IsPalindrome(x int) bool {
    reverse := 0
    tempX := x

    for x > 0 {
        mod := x % 10
        reverse = reverse * 10 + mod
        
        x /= 10
    }
    
    return reverse == tempX
}