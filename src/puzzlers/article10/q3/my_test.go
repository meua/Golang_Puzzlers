package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMyDemo(t *testing.T) {
	//myCh := make(chan int, 2)
	//defer close(myCh)
	//myStruct := MyStruct{"name:navy"}
	//myStruct.SendInt(myCh)
	//
	//intChan := getChanInt()
	//for elem := range intChan {
	//	fmt.Printf("The element in intchan2: %v\n", elem)
	//}

	ch := make(chan int, 2)
	//var ch chan int
	//ch <- 1
	close(ch)
	for i := range ch {
		println(i)
	}
}

func TestSwitchCase(t *testing.T) {
	intChannels := [...]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	for _, v := range intChannels {
		defer close(v)
	}
	//rand.NewSource(10)
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	default:
		fmt.Println("No candidate case is selected.")
	}
}

func TestCaseReturn(t *testing.T) {
	ch := make(chan int, 1)
	//ch <- 234
	time.AfterFunc(time.Second, func() {
		close(ch)
	})
	select {
	case v, ok := <-ch:
		if !ok {
			fmt.Printf("The candidate element is %d\n", v)
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}

type MyStruct struct {
	info string
}

func (t *MyStruct) SendInt(ch chan<- int) {
	ch <- rand.Int()
	fmt.Println("send ...")
}

func getChanInt() <-chan int {
	num := 5
	ch := make(chan int, num)
	defer close(ch)
	for i := 0; i < num; i++ {
		ch <- i
	}
	return ch
}
