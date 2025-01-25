package _390

import (
	"fmt"
)

func A() {
	defer w.Flush()

	As := readIntArray(r)

	prevA := As[0]
	count := 0
	isChangeable := true
	for i := 1; i < len(As); i++ {
		if prevA > As[i] {
			count++

			if prevA-As[i] > 1 {
				isChangeable = false
				break
			}
		}

		if count > 1 {
			isChangeable = false
			break
		}

		prevA = As[i]
	}

	if isChangeable && count == 1 {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
	}
}
