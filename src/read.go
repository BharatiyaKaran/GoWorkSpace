package main

import (
		"fmt"
		"os"
		"bufio"
		"strings"
)

const (
		maxlen=20
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	ret := make([]Name, 0, 1000)
	fmt.Println("Enter the file name:")
	fmt.Scan(&fileName)
	f, err:=os.Open(fileName)
	if err!=nil {
		fmt.Println("Error in opening file")
	}
	reader :=bufio.NewReader(f)
	var count int
	for {
		str, e :=reader.ReadString('\n')
		if e!=nil {
			break
		}
		str = strings.TrimSuffix(str, "\n")
		str2 := strings.Split(str, " ")
		//fmt.Println(str2)
		record := Name{fname:str2[0], lname:str2[1]}
		ret=append(ret, record)
		//fmt.Println(str)
		count++
	}

	for  _, r := range ret {
		fmt.Printf("First Name:%s, Last Name:%s\n", r.fname, r.lname)
	}
}


