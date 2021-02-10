package main

import (
	"fmt"
	"log"
	"math"
)

func GenDisplayFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
	}
	return fn
}

func main() {
	var acceleration float64
	fmt.Println("acceleration:")
	_, err := fmt.Scan(&acceleration)
	if err != nil {
		log.Fatal(err)
	}
	var initialVelocity float64
	fmt.Println("initial velocity:")
	_, err = fmt.Scan(&initialVelocity)
	if err != nil {
		log.Fatal(err)
	}
	var initialDisplacement float64
	fmt.Println("initial displacement:")
	_, err = fmt.Scan(&initialDisplacement)
	if err != nil {
		log.Fatal(err)
	}
	fn := GenDisplayFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
