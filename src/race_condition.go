package main

import "fmt"
import "sync"
import "time"

/* race condition occurs when 2 or more threads run for execution, 
   and the order of execution is unpredictable. 
   i.e we don't know which thread will win and execute the code block first.
   */

var wg sync.WaitGroup

func foo(i int) {
	fmt.Printf("I am thread %d\n", i)
	time.Sleep(time.Millisecond * 100) 
	wg.Done()
}

func main() {

	wg.Add(10)
	for i:=0; i<10; i++ {
		go foo(i)
	}
	wg.Wait()
}
