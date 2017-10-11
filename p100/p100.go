package main

import "fmt"

// This is the brute force way of solving this challenge

type Value struct {
	n     int
	count int
}

// importance of pointers
func (v *Value) Sequence() {
	if v.n > 1 {
		v.count = v.count + 1
		if v.n%2 == 0 {
			v.n = v.n / 2
			v.Sequence()
		} else {
			v.n = 3*v.n + 1
			v.Sequence()
		}
	}
}

func Series(i int, j int) {
	var x int

	x = i
	var maxcount = 0
	for x <= j {
		val := Value{
			n:     x,
			count: 1}
		val.Sequence()
		if maxcount < val.count {
			maxcount = val.count
		}
		x = x + 1
	}

	fmt.Println(i, j, maxcount)
}

func main() {
	// examples
	Series(1, 10)     // 20
	Series(100, 200)  // 125
	Series(201, 210)  // 89
	Series(900, 1000) // 525
}
