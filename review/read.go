package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which
contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and
print the first and last names found in each struct.
*/

type Person struct {
	fname string
	lname string
}

func parse(line string) *Person {
	var t []string
	var p Person
	if strings.Contains(line, " ") {
		t = strings.Split(line, " ")
		p.fname = t[0]
		p.lname = t[1]
	}
	return &p
}

func getNames(filePath string) *[]Person {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	s := make([]Person, 0)

	for scanner.Scan() {
		t := parse(scanner.Text())
		s = append(s, *t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &s
}

func printNames(people *[]Person) {
	for _, s := range *people {
		fmt.Println(s.fname + " " + s.lname)
	}
}

func main() {

	var filePath string

	fmt.Println("Enter location of file: ")
	_, err := fmt.Scanf("%s", &filePath)

	if err != nil {
		fmt.Println(err)
	}

	names := getNames(filePath)
	printNames(names)

}
