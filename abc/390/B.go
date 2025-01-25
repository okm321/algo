package _390

import "fmt"

func B() {
	defer w.Flush()

	N := readInt(r)
	As := readIntArray(r)

	prevA := As[0]
	var kohi float64
	isTohi := true
	for i := 1; i < N; i++ {
		if i == 1 {
			kohi = generateKohi(prevA, As[i])
			prevA = As[i]
			continue
		}

		currentKohi := generateKohi(prevA, As[i])
		if kohi != currentKohi {
			isTohi = false
			break
		}

		prevA = As[i]
	}

	if isTohi {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
	}
}

// 公比を求める
func generateKohi(prevInt, currentInt int) float64 {
	return float64(currentInt) / float64(prevInt)
}

// 正解
func B_another() {
	defer w.Flush()

	N := readInt(r)
	As := readIntArray(r)

	if N == 2 {
		fmt.Fprintln(w, "Yes")
	}

	for i := 1; i < N-1; i++ {
		a := As[i-1]
		b := As[i]
		c := As[i+1]

		if b*b != a*c {
			fmt.Fprintln(w, "No")
			return
		}
	}

	fmt.Fprintln(w, "Yes")
}
