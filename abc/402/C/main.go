package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()

	// N: N種類の食材
	// M: M個の料理
	N, M := read2Ints(r)

	// 料理の情報
	a := make([][]int, M)   // 各料理に含まれる食材のリスト
	idx := make([][]int, N) // 各食材が含まれる料理のリスト
	cnt := make([]int, M)   // 各料理に残っている苦手な食材の数

	// 料理の情報を読み込む
	for i := 0; i < M; i++ {
		input := readIntArray(r)
		k := input[0] // 料理iに使われる食材の数
		cnt[i] = k
		a[i] = make([]int, k)

		for j := 0; j < k; j++ {
			e := input[j+1] - 1 // 0-indexedに変換
			a[i][j] = e
			idx[e] = append(idx[e], i) // 食材eを使用している料理iを記録
		}
	}

	// 日毎に克服する食材
	B := readIntArray(r)

	// 克服カウント（累積）
	ans := 0

	// 各日について処理
	for i := 0; i < len(B); i++ {
		b := B[i] - 1 // 0-indexedに変換

		// 食材bを含む全ての料理を更新
		for _, id := range idx[b] {
			cnt[id]--
			if cnt[id] == 0 {
				ans++
			}
		}

		fmt.Fprintln(w, ans)
	}
}

// ── 数値読み取り ────────────────────────────────────────────────────

// readInt 単一の整数を読み取る関数
func readInt(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(line))
	return val
}

// readBigInt 単一の整数を読み取る関数
func readBigInt(reader *bufio.Reader) *big.Int {
	line, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(line))
	return big.NewInt(int64(val))
}

// read2Ints 一行に2つの整数のみの入力を読み込む
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])
	return i1, i2
}

// read3Ints 一行に3つの整数のみの入力を読み込む
func read3Ints(r *bufio.Reader) (int, int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])
	i3, _ := strconv.Atoi(strs[2])
	return i1, i2, i3
}

// readIntArray 整数の配列をよもとる関数
func readIntArray(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	arr := make([]int, len(parts))
	for i, s := range parts {
		arr[i], _ = strconv.Atoi(strings.TrimSpace(s))
	}

	return arr
}

// ── 文字列の読み取り ────────────────────────────────────────────────

// readString 単一の文字列を読み取る関数
func readString(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// readStringArray 文字列の配列を読み取る関数
func readStringArray(reader *bufio.Reader) []string {
	line, _ := reader.ReadString('\n')
	return strings.Fields(line)
}

// ── グリッド関係のやつ ──────────────────────────────────────────────

// readGrid height行の文字列グリッドを読み込む
//
// input:
// #..
// .#.
// ..#
//
// output:
// [
//
//	["#", ".", "."],
//	[".", "#", "."],
//	[".", ".", "#"]
//
// ]
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readString(r)
		grid[i] = strings.Split(str, "")
	}
	return grid
}

// readIntGrid height行の整数グリッドを読み込む
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArray(r)
	}

	return grid
}

// createGrid height行、width列のT型グリッドを作成
func createGrid[T any](height, width int, val T) [][]T {
	grid := make([][]T, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]T, width)
		for j := 0; j < width; j++ {
			grid[i][j] = val
		}
	}
	return grid
}

// writeGrid 文字列グリッドを出力する
func writeGrid(w *bufio.Writer, grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, strings.Join(grid[i], ""), "\n")
	}
}
