package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp/v3"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func gcdRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRecursive(b, a%b)
}

func main() {
	defer w.Flush()

	N := readInt(r)

	prereqs := make([][]int, N+1)
	unlock := make([][]int, N+1)
	skilled := make([]bool, N+1)

	queue := []int{}

	pp.Println(gcdRecursive(88, 33))

	for i := 1; i <= N; i++ {
		A, B := read2Ints(r)
		prereqs[i] = []int{A, B}

		if A == 0 && B == 0 {
			queue = append(queue, i)
			skilled[i] = true
		} else {
			if A != 0 {
				unlock[A] = append(unlock[A], i)
			}
			if B != 0 {
				unlock[B] = append(unlock[B], i)
			}
		}
	}

	// BFS
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, nextSkill := range unlock[current] {
			if skilled[nextSkill] {
				continue
			}

			A, B := prereqs[nextSkill][0], prereqs[nextSkill][1]
			if skilled[A] || skilled[B] {
				skilled[nextSkill] = true
				queue = append(queue, nextSkill)
			}
		}
	}

	count := 0
	for i := 1; i <= N; i++ {
		if skilled[i] {
			count++
		}
	}

	fmt.Fprintln(w, count)
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

// readStringSplit 文字列を指定のセパレータで分割して読み取る関数
func readStringSplit(reader *bufio.Reader, sep string) []string {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return strings.Split(line, sep)
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

func convertIntArrToString(arr []int) string {
	strSlice := make([]string, len(arr))
	for i, v := range arr {
		strSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(strSlice, " ")
}
