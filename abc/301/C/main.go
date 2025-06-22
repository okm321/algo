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

func count(ss string) map[string]int {
	m := make(map[string]int, 27)

	for _, s := range strings.Split(ss, "") {
		if s == "@" {
			m["@"]++
		} else {
			m[s]++
		}
	}
	return m
}

func main() {
	defer w.Flush()

	S := readString(r)
	T := readString(r)
	cs := count(S)
	ct := count(T)

	swapMap := map[string]struct{}{
		"a": {},
		"t": {},
		"c": {},
		"o": {},
		"d": {},
		"e": {},
		"r": {},
	}

	for k := range cs {
		// 同じ文字を間引く
		if _, ok := ct[k]; ok && k != "@" {
			min := cs[k]
			if ct[k] < min {
				min = ct[k]
			}
			cs[k] -= min
			ct[k] -= min
			if cs[k] == 0 {
				delete(cs, k)
			}
			if ct[k] == 0 {
				delete(ct, k)
			}
		}
	}

	for k := range cs {
		if _, ok := swapMap[k]; !ok && k != "@" {
			fmt.Fprintln(w, "No")
			return
		}
	}
	for k := range ct {
		if _, ok := swapMap[k]; !ok && k != "@" {
			fmt.Fprintln(w, "No")
			return
		}
	}

	csSwappableCount := cs["@"]
	ctSwappableCount := ct["@"]
	csUnSwappableCount := 0
	ctUnSwappableCount := 0
	for k, v := range cs {
		if k == "@" {
			continue
		}
		csUnSwappableCount += v
	}
	for k, v := range ct {
		if k == "@" {
			continue
		}
		ctUnSwappableCount += v
	}

	if csSwappableCount < ctUnSwappableCount || ctSwappableCount < csUnSwappableCount {
		fmt.Fprintln(w, "No")
		return
	}

	fmt.Fprintln(w, "Yes")
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
