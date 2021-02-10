package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(sli []int) {
	for {
		swapped := false
		for i := 1; i < len(sli); i++ {
			if sli[i-1] > sli[i] {
				Swap(sli, i-1)
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}
}

func Swap(sli []int, pos int) {
	tmp := sli[pos]
	sli[pos] = sli[pos+1]
	sli[pos+1] = tmp
}

func stoiSlices(str_slice []string) []int {
	int_slice := []int{}

	for _, elem := range str_slice {
		i, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		int_slice = append(int_slice, i)
	}

	return int_slice
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a list of less than 10 numbers: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str_slice := strings.Fields(str)
	if len(str_slice) > 10 {
		log.Fatal("You enter more than 10 elements for your array")
	}

	var int_slice = stoiSlices(str_slice)

	BubbleSort(int_slice)

	for i := range int_slice {
		fmt.Printf("%d ", int_slice[i])
	}
}
