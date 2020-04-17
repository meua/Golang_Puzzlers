package main

import (
	"fmt"
	"testing"
)

func TestSwtich(t *testing.T) {
	values := [...]int{0, 1, 2, 3, 4, 5, 6}
	switch values[4] {
	case values[0], values[1], values[2]:
		fmt.Println("0 or 1")
	case values[2], values[3], values[4]:
		fmt.Println("2 or 3")
		//fallthrough
	case values[4], values[5], values[6]:
		fmt.Println("4 or 5 or 6")
	}
}

func TestSwtich0(t *testing.T) {
	values := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch values[2] {
	case 0, 1: //,2:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
		fallthrough
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}

func TestSwitch1(t *testing.T) {
	value6 := interface{}(byte(127))
	switch value6.(type) {
	case uint8, uint16, int32:
		fmt.Println("uint8 or uint16")
	//case byte, rune:
	//	fmt.Println("byte")
	default:
		fmt.Println("default")
	}
}
