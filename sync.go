package main

import (
	"strconv"
	"strings"
	"sync"
)

type ZeroEvenOddSync struct {
	N int

	output chan int
	mu     sync.Mutex
	i      int
	c      *sync.Cond
	wg     sync.WaitGroup
}

func NewZeroEvenOddSync(N int) ZeroEvenOdd {
	z := ZeroEvenOddSync{
		N: N,
	}
	z.c = sync.NewCond(&z.mu)
	return &z
}

func (z *ZeroEvenOddSync) String() string {
	if z.N == 0 {
		return ""
	}

	z.output = make(chan int, 5)

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

func (z *ZeroEvenOddSync) zero() {
	//prints 0
	defer z.wg.Done()

	for i := 0; i < z.N; i++ {
		z.c.L.Lock()
		for z.i != 0 {
			z.c.Wait()
		}

		z.output <- 0

		if i%2 == 0 {
			z.i = 1
		} else {
			z.i = 2
		}

		z.c.Broadcast()
		z.c.L.Unlock()
	}
}

func (z *ZeroEvenOddSync) even() {
	//print even numbers
	defer z.wg.Done()

	for i := 2; i <= z.N; i += 2 {
		z.c.L.Lock()
		for z.i != 2 {
			z.c.Wait()
		}

		z.output <- i
		z.i = 0

		z.c.Broadcast()
		z.c.L.Unlock()
	}
}

func (z *ZeroEvenOddSync) odd() {
	//print odd numbers
	defer z.wg.Done()

	for i := 1; i <= z.N; i += 2 {
		z.c.L.Lock()
		for z.i != 1 {
			z.c.Wait()
		}

		z.output <- i
		z.i = 0

		z.c.Broadcast()
		z.c.L.Unlock()
	}
}
