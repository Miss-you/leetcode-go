package main

import "fmt"

func canWinNim(n int) bool{
	if (n%4 == 0){
		return false
	} else{
		return true
	}
}

func main(){
	flag := canWinNim(4)
	fmt.Println("flag =", flag)
	
	flag = canWinNim(7)
	fmt.Println("flag =", flag)
}
