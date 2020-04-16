package main

import (
	"fmt"
	"testing"
)

func TestSeq(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
		go func() { fmt.Println(<-ch) }()
	}
	for {
		if len(ch) == 0 {
			break
		}
	}
}

func TestArrSum(t *testing.T) {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func TestSliceSum(t *testing.T) {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}
