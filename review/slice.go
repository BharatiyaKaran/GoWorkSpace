package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an
empty integer slice of size (length) 3. During each pass through the loop, the program prompts the
user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts
the slice, and prints the contents of the slice in sorted order. The slice must grow in size to
accommodate any number of integers which the user decides to enter. The program should only quit
(exiting the loop) when the user enters the character â€˜Xâ€™ instead of an integer.
*/

func main() {
	var s1 string
	num := make([]float64, 0)

	for {

		fmt.Println("Enter a number : (x to exit)")
		_, err := fmt.Scanf("%s", &s1)

		if err != nil {
			fmt.Println("Error:")
			fmt.Println(err)
			fmt.Println("\nSlice values:")
			fmt.Println(num)
			fmt.Println()
			continue
		}

		if strings.ToLower(s1) == "x" {
			break
		}

		d, err := strconv.ParseFloat(s1, 64)
		if err == nil {
			num = append(num, float64(int64(d)))
		}
		sort.Float64s(num)
		fmt.Println("\nSlice values:")
		fmt.Println(num)
		fmt.Println()
	}
}
