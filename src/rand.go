package main

import (
		"fmt"
		//"sync"
		"time"
		"math/rand"
)



func main(){
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	for i:=0; i<20; i++ {
		fmt.Println(r1.Intn(20))
	}
}
