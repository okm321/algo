package _389

import (
	"fmt"
	"math/big"
)

func B() {
	defer w.Flush()

	X := readBigInt(r)

	fmt.Fprintln(w, reverseKaijou(X))
}

// 6を入れると3が帰ってくるやつ
func reverseKaijou(n *big.Int) *big.Int {
	k := big.NewInt(1)
	current := big.NewInt(1)

	for current.Cmp(n) < 0 {
		k.Add(k, big.NewInt(1))
		current.Mul(current, k)
	}

	if current.Cmp(n) == 0 {
		return k
	}
	return nil
}
