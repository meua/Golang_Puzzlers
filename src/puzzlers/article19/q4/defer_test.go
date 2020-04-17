package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	/*	for i := 0; i < 6 ; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}*/
	fmt.Println("Enter function main")
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("defer exit")
		}
	}()
	defer panic(errors.New("defer panic"))
	fmt.Println("Exit function main")
}

func TestDeferReturn(t *testing.T) {
	println(deferReturn())
}

func deferReturn() (i int) {
	i = 0
	defer func(i int) {
		i += 1
	}(i)
	i += 1
	return
}
