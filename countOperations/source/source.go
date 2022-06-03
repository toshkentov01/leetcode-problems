package source

func CountOperations(num1 int, num2 int) int {
    counter := 1
    
    if num1 == 0 && num2 == 0 || num1 == 0 || num2 == 0 {
        return 0
    }
    
    for num1 != 0 {
        if num1 > num2 {
            num1 = num1 - num2
            counter += 1
        
        } else if num1 < num2 {
            num2 = num2 - num1
            counter += 1
        
        } else if num1 == num2 {
            return counter 
        }
    }   
    
    return counter
}