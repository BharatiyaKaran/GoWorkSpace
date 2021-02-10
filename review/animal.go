package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name   string
	food   string
	motion string
	sound  string
}

func (a *Animal) SetMotion(motion string) {
	a.motion = motion
}

func (a Animal) GetMotion() string {
	return a.motion
}

func (a *Animal) SetSound(sound string) {
	a.sound = sound
}

func (a Animal) GetSound() string {
	return a.sound
}

func (a *Animal) SetFood(food string) {
	a.food = food
}

func (a Animal) GetFood() string {
	return a.food
}

func (a *Animal) SetName(name string) {
	a.name = name
}

func (a Animal) GetName() string {
	return a.name
}

func (a Animal) Move() {
	fmt.Println(a.GetName() + ", locomotion method: " + a.GetMotion())
}

func (a Animal) Speak() {
	fmt.Println(a.GetName() + ", spoken sound: " + a.GetSound())
}

func (a Animal) Eat() {
	fmt.Println(a.GetName() + ", food eaten: " + a.GetFood())
}

func main() {

	var line string
	var input []string
	var animals [3]Animal
	animals[0].SetFood("grass")
	animals[0].SetSound("moo")
	animals[0].SetMotion("walk")
	animals[0].SetName("cow")

	animals[1].SetFood("worms")
	animals[1].SetSound("peep")
	animals[1].SetMotion("fly")
	animals[1].SetName("bird")

	animals[2].SetFood("mice")
	animals[2].SetSound("hsss")
	animals[2].SetMotion("slither")
	animals[2].SetName("snake")

	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Print("> ")
		line = scanner.Text()
		input = strings.Fields(line)
		for _, x := range animals {
			if strings.ToLower(x.GetName()) == input[0] {
				switch strings.ToLower(input[1]) {
				case "eat":
					{
						x.Eat()
					}
				case "move":
					{
						x.Move()
					}
				case "speak":
					{
						x.Speak()
					}
				}
				break
			}
		}
	}
}
