func addDigits(num int) int {
    if (num == 0){
        return 0
    }
    
    finalDigit := num%9
    if (finalDigit == 0){
        return 9
    } else{
        return finalDigit
    }
}
