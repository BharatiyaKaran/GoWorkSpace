package main

import "fmt"

func main() {
	var num float32
	var truncNum int
	fmt.Printf("Enter a float number:")
	fmt.Scan(&num)
	truncNum = int(num)

	fmt.Printf("the truncated number is %d\n", truncNum)

	fmt.Printf("Enter a float number:")
	fmt.Scan(&num)
	truncNum = int(num)

	fmt.Printf("the truncated number is %d\n", truncNum)

}
