package source

func isPowerOfThree(n int) bool {    
    
    if mod := n % 3; mod == 0 && n != 0 {
		return isPowerOfThree(n/3)
	}
    
    return (n==3 || n == 1)
}