package main

import "fmt"

// function to print the pattern
func fact(n int){
	fmt.Println("Pattern :")
	for i:=1;i<=n;i++{
		var num=i
		for j:=1;j<=i;j++{
		fmt.Print(" ",num)
		num=num+n-j;

	}
	fmt.Println()
}
}


func main(){
	var num int
	fmt.Println("Enter the Number :")
	fmt.Scanf("%d",&num)
	fact(num)

}