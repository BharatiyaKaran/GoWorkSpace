package main

import "fmt"
import "strings"

func main() {
	var str1 string
	fmt.Scan(&str1)
	if strings.Contains(str1, "i") && 
		strings.Contains(str1, "a") &&
		strings.Contains(str1, "n") {
			fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
