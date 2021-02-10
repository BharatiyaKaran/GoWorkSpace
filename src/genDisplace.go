package main

import "fmt"

func GenDisplaceFn(acc, v0, s0 float64) func (float64) float64 {
	return func(t float64) float64 {
	disp:= (0.5)*acc*t*t + v0*t + s0
	return disp
	}
}

func main(){
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement")
	var acc, v0, s0, t float64
	fmt.Scan(&acc, &v0, &s0)
	f1:=GenDisplaceFn(acc, v0, s0)
	fmt.Println("Enter value of time")
	fmt.Scan(&t)
	fmt.Printf("Final displacement is=%f\n", f1(t))
}
