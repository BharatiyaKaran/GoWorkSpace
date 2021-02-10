package main

import "fmt"
import "strings"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
	food string
	loco string
	speak string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.loco)
}
func (c Cow) Speak() {
	fmt.Println(c.speak)
}

type Bird struct {
	name string
	food string
	loco string
	speak string
}
func (b Bird) Eat() {
	fmt.Println(b.food)
}
func (b Bird) Move() {
	fmt.Println(b.loco)
}
func (b Bird) Speak() {
	fmt.Println(b.speak)
}

type Snake struct {
	name string
	food string
	loco string
	speak string
}
func (s Snake) Eat() {
	fmt.Println(s.food)
}
func (s Snake) Move() {
	fmt.Println(s.loco)
}
func (s Snake) Speak() {
	fmt.Println(s.speak)
}

func main() {
	var str1, str2, str3 string
	Amap:=make(map[string]Animal)
	for {
		fmt.Print("> ")
		fmt.Scan(&str1, &str2, &str3)
		if str1=="end" || str1=="exit" {
			break
		}
		if strings.ToLower(str1)=="newanimal" {
			var Anew Animal
			if strings.ToLower(str3)=="cow" {
				Anew = Cow{str2, "grass", "walk", "moo"}
				Amap[str2]=Anew
			}
			if strings.ToLower(str3)=="bird" {
				Anew = Bird{str2, "worms", "fly", "peep"}
				Amap[str2]=Anew
			}
			if strings.ToLower(str3)=="snake" {
				Anew = Snake{str2, "mice", "slither", "hsss"}
				Amap[str2]=Anew
			}
			fmt.Println("Created it!")
		}

		if strings.ToLower(str1)=="query" {
			if strings.ToLower(str3)=="eat" {
                if Amap[str2]!=nil {
					Amap[str2].Eat()
				}
            }
            if strings.ToLower(str3)=="move" {
                if Amap[str2]!=nil {
					Amap[str2].Move()
				}
            }
            if strings.ToLower(str3)=="speak" {
                if Amap[str2]!=nil {
					Amap[str2].Speak()
				}
            }
		}
	}
}
