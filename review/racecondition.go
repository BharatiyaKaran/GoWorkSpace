/* Race condition occurs due to the two functions access the variable called ‘total’ concurrently.
 *This causes the incorrect total and average to be displayed to the user.
 */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sliceTotal(sli []int, total *int) {
	//for i := 0; i < len(sli); i++ {
	for i := 0; i < 1000000; i++ {
		//*total += sli[i]
		*total += i
	}
	wg.Done()
}

func average(length int, total *int, avg *float64) {
	tot := float64(*total)
	//ln := float64(length)
	//*avg = tot / ln
	*avg = tot / 1000000
	wg.Done()
}

func main() {
	var total int
	var avg float64
	sli := []int{1, 2, 3, 4, 5,6,7,8,9}

	wg.Add(2)
	go sliceTotal(sli, &total)
	go average(len(sli), &total, &avg)
	//time.Sleep(time.Second)
	fmt.Println("Total integers in slice: ", total)
	fmt.Println("Average: ", avg)
	wg.Wait()
}
