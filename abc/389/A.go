package _389

import (
	"fmt"
	"strconv"
)

func A() {
	defer w.Flush()

	S := readString(r)

	first := S[0]
	third := S[2]

	intFirst, err := strconv.Atoi(string(first))
	if err != nil {
		panic(err)
	}
	intThird, err := strconv.Atoi(string(third))
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, intFirst*intThird)
}
