package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	fmt.Println("Enter function main")
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit function defer")
	}()
	panic(errors.New("something error"))
	fmt.Println("Exit function main")
}
