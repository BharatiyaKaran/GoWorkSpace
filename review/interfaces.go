package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Bird) Eat() {
	fmt.Println("worms")
}

func (c Bird) Move() {
	fmt.Println("fly")
}

func (c Bird) Speak() {
	fmt.Println("peep")
}

func (c Snake) Eat() {
	fmt.Println("mice")
}

func (c Snake) Move() {
	fmt.Println("slither")
}

func (c Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := map[string]Animal{}

	for {
		arg1, arg2, arg3 := readInput("Type your request (e.g newanimal John cow)")

		if arg1 == "newanimal" {
			// Create animal
			switch arg3 {
			case "cow":
				animals[arg2] = new(Cow)
				break
			case "bird":
				animals[arg2] = new(Bird)
				break
			case "snake":
				animals[arg2] = new(Snake)
			default:
				log.Fatal("Unknown animal type")
			}
			fmt.Println("Created it!")
		} else if arg1 == "query" {
			// Query
			var animal = animals[arg2]
			switch arg3 {
			case "eat":
				animal.Eat()
				break
			case "move":
				animal.Move()
				break
			case "speak":
				animal.Speak()
			default:
				log.Fatal("Unknown animal method")
			}
		} else {
			log.Fatal("Unknown request")
		}
	}

}

func readInput(question string) (string, string, string) {
	fmt.Println(question)
	fmt.Print(">")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	text := strings.Split(scanner.Text(), " ")

	return text[0], text[1], text[2]
}
