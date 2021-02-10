package main

import (
	"fmt"
	"log"
	"strings"
)


func main() {
	var userInput string
	prefix:= "i"
	suffix := "n"

	fmt.Println("Input a string: ")
	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		log.Fatal(err)
	}

	word := strings.ToLower(userInput)

	if strings.HasPrefix(word, prefix) {

		if strings.Contains(word, "a") {
		} else {
			fmt.Println("Not Found!")
		}

			if strings.HasSuffix(word, suffix) {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}

	} else {
		fmt.Println("Not Found!")
	}
}