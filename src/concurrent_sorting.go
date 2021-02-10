package main

import (
		"fmt"
		"sync"
		"sort"
		//"math/rand"
		//"time"
)

var wg sync.WaitGroup

// Sorting function
func sortFun(nums []int, sorted chan []int) {
	fmt.Printf("Sorting:%v\n", nums)
	sort.Ints(nums)
	sorted <- nums
	wg.Done()
}


// Merge function
func merge(nums1, nums2 []int) []int {
	left:=len(nums1)
	right:=len(nums2)
	var total int = left+right
	merged := make([]int, total)
	i:=0
	j:=0
	k:=0
	for i<left && j<right {
		if nums1[i]< nums2[j] {
			merged[k]=nums1[i]
			i++
			k++
		} else {
			merged[k]=nums2[j]
			j++
			k++
		}
	}

	for i<left {
		merged[k]=nums1[i]
		i++
		k++
	}

	for j<right {
		merged[k]=nums2[j]
        j++
        k++
	}
	return merged
}

func main(){
	var num int
	var n int
	//s1 := rand.NewSource(time.Now().UnixNano())
    //r1 := rand.New(s1)
	fmt.Println("Enter the number of numbers you wish to sort")
	fmt.Scan(&n)
	nums := make([]int, 0, 100)
	fmt.Println("Enter the numbers to sort")
	for i:=0; i<n; i++ {
		fmt.Scan(&num)
		//num=r1.Intn(20)
		nums=append(nums, num)
	}
	// get the quotient for dividing into 4 threads
	k:=n/4
	// make a channels
	done:=make(chan []int)
	for i:=0; i<3; i++ {
		wg.Add(1)
		go sortFun(nums[k*i: k*(i+1)], done)
	}
	wg.Add(1)
	go sortFun(nums[3*k: ], done)
	var mergedVal = make([]int, 0)
	count:=0
	for sortedVal:= range done {
		mergedVal=merge(mergedVal, sortedVal)
		count++
		if count==4 {
			break
		}
	}
	fmt.Println("After sorting...")
	fmt.Println(mergedVal)
	wg.Wait()
}
