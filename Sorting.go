package main

import "fmt"

// function for sorting array using bubble sort
func sort(arr []int64,n int)  []int64 {
	for i:=0;i<n-1;i++{
		for j:=0;j<n-i-1;j++{

			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
    return arr
}

func main(){
	var size int
	fmt.Print("Enter the Size Of Array :")
	fmt.Scanf("%d",&size)
	input := make([]int64, size)
	for i:=0;i<size;i++{
		fmt.Scan(&input[i])
	}
	fmt.Println("Array Before Sorting :",input)
	fmt.Println("Array After Sorting :",sort(input,size))
}