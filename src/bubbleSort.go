package main

import "fmt"

func Swap(a, b *int) {
	temp := *a
	*a=*b
	*b=temp
}

func BubbleSort(numSlice []int) {
	n:=len(numSlice)
	for i:=0; i<n-1; i++ {
		for j:=0; j<n-i-1; j++ {
			if numSlice[j] > numSlice[j+1] {
				Swap(&numSlice[j], &numSlice[j+1])
			}
		}
	}
}

func main(){
	fmt.Println("Enter 10 numbers")
	var num int
	numSlice := make([]int, 0, 10)
	for i:=0; i<10; i++ {
		fmt.Scan(&num)
		numSlice=append(numSlice, num)
	}
	fmt.Println("You entered: ")
	fmt.Println(numSlice)
	BubbleSort(numSlice)
	fmt.Println("After Bubble Sort")
	fmt.Println(numSlice)
}
