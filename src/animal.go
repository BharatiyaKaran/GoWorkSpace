package main

import "fmt"
import "strings"

type Animal struct {
	food 	string
	loco 	string
	noise 	string
}

func (A Animal) eat() string {
	return A.food
}

func (A Animal) Move() string {
	return A.loco
}

func (A Animal) Speak() string {
	return A.noise
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{food:"worms", loco:"fly", noise:"peep"}
	snake := Animal{food:"mice", loco:"slither", noise:"hsss"}
	var str1, str2 string
	for {
		fmt.Print("> ")
		fmt.Scan(&str1, &str2)
		if str1=="end" || str1=="exit" {
			break
		}
		if strings.ToLower(str1)=="cow" {
			if strings.ToLower(str2)=="eat" {
				fmt.Println(cow.eat())
			}
			if strings.ToLower(str2)=="move" {
				fmt.Println(cow.Move())
			}
			if strings.ToLower(str2)=="speak" {
				fmt.Println(cow.Speak())
			}
		}
		if strings.ToLower(str1)=="bird" {
			if strings.ToLower(str2)=="eat" {
				fmt.Println(bird.eat())
			}
			if strings.ToLower(str2)=="move" {
				fmt.Println(bird.Move())
			}
			if strings.ToLower(str2)=="speak" {
				fmt.Println(bird.Speak())
			}
		}
		if strings.ToLower(str1)=="snake" {
			if strings.ToLower(str2)=="eat" {
				fmt.Println(snake.eat())
			}
			if strings.ToLower(str2)=="move" {
				fmt.Println(snake.Move())
			}
			if strings.ToLower(str2)=="speak" {
				fmt.Println(snake.Speak())
			}
		}
	}
}
