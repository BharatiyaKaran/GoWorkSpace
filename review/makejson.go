package main

import (
	"encoding/json"
	"fmt"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the
keys â€œnameâ€ and â€œaddressâ€, respectively. Your program should use Marshal() to create a
JSON object from the map, and then your program should print the JSON object.
*/

type Person struct {
	Name    string
	Address string
}

func main() {
	var person Person

	fmt.Println("Enter a name: ")
	_, err := fmt.Scanf("%s", &person.Name)

	if err != nil {
		fmt.Println(err)
	}
	err = nil

	fmt.Println("Enter an address: ")
	_, err = fmt.Scanf("%s", &person.Address)

	if err != nil {
		fmt.Println(err)
	}

	s, er := json.Marshal(person)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(s))

}
