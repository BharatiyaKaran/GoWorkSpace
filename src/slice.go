package main

import "fmt"
import "sort"
import "strconv"
import "strings"

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an
empty integer slice of size (length) 3. During each pass through the loop, the program prompts the
user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts
the slice, and prints the contents of the slice in sorted order. The slice must grow in size to
accommodate any number of integers which the user decides to enter. The program should only quit
(exiting the loop) when the user enters the character "X" instead of an integer.
*/


func main() {
    	var x string
    	var intSlice = make([]int, 0, 1000)
	for {
    	fmt.Println("Enter a number")
    	fmt.Scan(&x)
		z, err := strconv.ParseInt(x, 10, 32)
		if err==nil {
    		intSlice=append(intSlice, int(z))
		}
    	fmt.Println("Sorted output")
    	sort.Ints(intSlice)
    	fmt.Println(intSlice)
		if strings.ToLower(x)=="x" {
			fmt.Println("Exiting...")
			break
		}
    }
}
