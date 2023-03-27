package main

import (
	"fmt"
)

//the goal is to print concarrently 010203040506070809......0

type ZeroEvenOdd interface {
	String() string
}

func main() {
	z := NewZeroEvenOddAtomic(0)
	fmt.Printf("%s\n", z.String())
	fmt.Println("-----------------")
	z = NewZeroEvenOddAtomic(1)
	fmt.Printf("%s\n", z.String())
	fmt.Println("-----------------")
	z = NewZeroEvenOddAtomic(2)
	fmt.Printf("%s\n", z.String())
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	z = NewZeroEvenOddSync(0)
	fmt.Printf("%s\n", z.String())
	fmt.Println("-----------------")
	z = NewZeroEvenOddSync(1)
	fmt.Printf("%s\n", z.String())
	fmt.Println("-----------------")
	z = NewZeroEvenOddSync(2)
	fmt.Printf("%s\n", z.String())
}
