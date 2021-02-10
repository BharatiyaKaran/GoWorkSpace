package main

import(
	"fmt"
	"encoding/json"
)

type record struct {
	Name string	`json:"name"`
	Addr string	`json:"address"`
}

func main() {
	var name1, addr1 string
	fmt.Println("Enter name and address")
	fmt.Scan(&name1, &addr1)
	r1 := record{Name: name1, Addr:addr1}
	jStr, err := json.Marshal(r1)
	if err!=nil {
		fmt.Println("Some error occured")
	}
	fmt.Println(string(jStr))
}
