package main

import "fmt"

// function calculating factorial
func fact(n int) int{
	if n > 1 {
        return n*fact(n-1)
    } else {
		return 1
	}

}

func main(){
	var num int
	fmt.Println("Enter the Number :")
	fmt.Scanf("%d",&num)
	fmt.Println("Factorial of number",num,"is :",fact(num))

}