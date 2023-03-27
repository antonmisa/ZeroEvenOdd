package main

import (
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type ZeroEvenOddAtomic struct {
	N int

	output chan int
	v      atomic.Uint32
	wg     sync.WaitGroup
}

func NewZeroEvenOddAtomic(N int) ZeroEvenOdd {
	z := ZeroEvenOddAtomic{
		N: N,
	}
	return &z
}

func (z *ZeroEvenOddAtomic) String() string {
	if z.N == 0 {
		return ""
	}

	z.output = make(chan int, 5)
	z.v.Store(0)

	z.wg.Add(3)
	go z.zero()
	go z.even()
	go z.odd()
	go func() {
		z.wg.Wait()
		close(z.output)
	}()

	result := strings.Builder{}
	result.Grow(2 * z.N)
	for data := range z.output {
		result.WriteString(strconv.Itoa(data))
	}

	return result.String()
}

func (z *ZeroEvenOddAtomic) zero() {
	//prints 0
	defer z.wg.Done()

	for i := 0; i < z.N; i++ {
		for {
			if z.v.CompareAndSwap(0, 100) {
				break
			}
		}

		z.output <- 0

		if i%2 == 0 {
			z.v.Store(1)
		} else {
			z.v.Store(2)
		}
	}
}

func (z *ZeroEvenOddAtomic) even() {
	//print even numbers
	defer z.wg.Done()

	for i := 2; i <= z.N; i += 2 {
		for {
			if z.v.CompareAndSwap(2, 20) {
				break
			}
		}

		z.output <- i
		z.v.Store(0)
	}
}

func (z *ZeroEvenOddAtomic) odd() {
	//print odd numbers
	defer z.wg.Done()

	for i := 1; i <= z.N; i += 2 {
		for {
			if z.v.CompareAndSwap(1, 10) {
				break
			}
		}

		z.output <- i
		z.v.Store(0)
	}
}
