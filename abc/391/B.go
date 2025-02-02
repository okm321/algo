package _391

import (
	"fmt"
)

func B() {
	N, M := read2Ints(r)

	S := readGrid(r, N)
	T := readGrid(r, M)

	for a := 1; a < N+1; a++ {
		for b := 1; b < N+1; b++ {
			match := true
			for i := 0; i < M; i++ {
				for j := 0; j < M; j++ {
					if S[a+i][b+j] != T[i][j] {
						match = false
					}
				}
			}
			if match {
				fmt.Fprintln(w, a+1, b+1)
				return
			}
		}
	}
}
