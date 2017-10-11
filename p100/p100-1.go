package main

import "fmt"

// Note: below concept is not quite right, but this on track to create a better
//		than brute-force.

// The sequences of the cycles are repeating patterns for example:
// 		9 28 14 7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1
//				  22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1
// A faster solution than brute force would be to keep the sequence
//		of resulting numbers, if an integer is in this sequence then
//		it's cycle will be shorter than the one already evaluated.
//		If it is not in the sequence, then run the algorithm on it.
// Also, you could run the algorithm until you reach the integer
//		which started the last longest cycle. Adding the cycle length
//		thus far and the old one will result in the new length

type Value2 struct {
	n     int
	seq   []int
	lgseq []int
}

func (v *Value2) Sequence2() {
	v.seq = append(v.seq, v.n)
	if v.n == 1 {
		v.lgseq = v.seq
	} else if len(v.lgseq) > 0 && v.lgseq[0] == v.n {
		fmt.Println("reached beginning of longest sequence", v.seq)
		v.lgseq = append(v.seq, v.lgseq[1:]...)
	} else if v.n > 1 {
		fmt.Println("current sequence: ", v.seq)
		if v.n%2 == 0 {
			v.n = v.n / 2
			v.Sequence2()
		} else {
			v.n = 3*v.n + 1
			v.Sequence2()
		}
	}
}

func (v *Value2) Series2(i int, j int) {
	x := i

	for x <= j {
		inlgseq := false
		for _, l := range v.lgseq {
			if x == l {
				inlgseq = true
			}
		}
		fmt.Println(inlgseq)
		if !(inlgseq) {
			v.n = x
			v.seq = []int{}
			fmt.Println("n: ", v.n)
			v.Sequence2()
		}
		fmt.Println("longest sequence: ", v.lgseq)
		x = x + 1
	}
}

// fancy try
// val := Value2{
// 	n:     0,
// 	seq:   make([]int, 0),
// 	lgseq: make([]int, 0),
// }

// val.Series2(1, 10)
// fmt.Println("result: ", val.lgseq)
