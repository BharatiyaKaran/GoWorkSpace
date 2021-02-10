package main

import "fmt"

func modifySlice(s []int) {
	for i,_:=range s {
		s[i]++
	}
}

func main() {
	s1 := []int{1,2,3,4,5}
	fmt.Println(s1)
	modifySlice(s1)
	fmt.Println(s1)
}
