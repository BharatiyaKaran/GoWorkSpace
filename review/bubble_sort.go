package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	tmp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = tmp
}

func main() {
	slice := make([]int, 0, 0)

	fmt.Print("Input (seperated by space): ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		tokens := strings.Fields(input)

		for _, token := range tokens {
			val, _ := strconv.Atoi(token)
			slice = append(slice, val)
		}

		BubbleSort(slice)

		fmt.Println(slice)
	}
}
