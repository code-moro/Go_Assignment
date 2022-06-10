package main

import  "fmt"

// function for reversing the string
func Reverse(n string)string{
	b:=[]byte(n)
	for i:=0;i<len(b)/2;i++{
		b[len(b)-i-1],b[i]=b[i],b[len(b)-i-1]
	}
	return string(b)
}

func main(){
	var str string
	fmt.Println("Enter the String:")
    fmt.Scanf("%s", &str)
	fmt.Println("Reversed String is:")
	fmt.Print(Reverse(str))
}