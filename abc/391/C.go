package _391

import (
	"fmt"
)

func C() {
	N, Q := read2Ints(r)

	var qs [][]int
	for i := 0; i < Q; i++ {
		qs = append(qs, readIntArray(r))
	}

	birds := make(map[int]int)
	for i := 0; i < N; i++ {
		birds[i] = 0
	}

	for i := 0; i < len(qs); i++ {
		if len(qs[i]) == 1 {
			fmt.Fprintln(w, birds)
			continue
		}

		birds[qs[i][2]] += birds[qs[i][1]]
		birds[qs[i][1]] = 0
	}

}
