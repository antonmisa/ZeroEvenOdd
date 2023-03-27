package main

import (
	"strconv"
	"testing"
)

var tests = []struct {
	N        int
	Expected string
}{
	{0, ""},
	{1, "01"},
	{2, "0102"},
	{10, "010203040506070809010"},
}

func TestZeroEvenOddAtomic(t *testing.T) {
	for _, test := range tests {
		z := NewZeroEvenOddAtomic(test.N)
		v := z.String()
		if v != test.Expected {
			t.Fatalf("For %d expected %s, but found %s", test.N, test.Expected, v)
		}
	}
}

func TestZeroEvenOddSync(t *testing.T) {
	for _, test := range tests {
		z := NewZeroEvenOddSync(test.N)
		v := z.String()
		if v != test.Expected {
			t.Fatalf("For %d expected %s, but found %s", test.N, test.Expected, v)
		}
	}
}

var benchmarks = []int{100, 1000, 10000, 100000}

func BenchmarkZeroEvenOddAtomic(t *testing.B) {
	for _, bench := range benchmarks {
		t.Run("Length"+strconv.Itoa(bench), func(t *testing.B) {
			z := NewZeroEvenOddAtomic(bench)
			z.String()
		})
	}
}

func BenchmarkZeroEvenOddSync(t *testing.B) {
	for _, bench := range benchmarks {
		t.Run("Length"+strconv.Itoa(bench), func(t *testing.B) {
			z := NewZeroEvenOddSync(bench)
			z.String()
		})
	}
}
