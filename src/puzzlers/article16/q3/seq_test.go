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
