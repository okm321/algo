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

	H, W := read2Ints(r)

	grid := readGrid(r, H)

	targetStr := "snuke"
	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for dir := 0; dir < 8; dir++ {
				found := true
				// 1文字ずつチェック
				for t := 0; t < len(targetStr); t++ {
					ni := i + t*di[dir]
					nj := j + t*dj[dir]
					
					// 範囲外チェック
					if ni < 0 || ni >= H || nj < 0 || nj >= W {
						found = false
						break
					}
					
					// 文字が違ったら即座に中断
					if grid[ni][nj] != string(targetStr[t]) {
						found = false
						break
					}
				}
				
				// "snuke"が見つかったら
				if found {
					// 5マスの座標を出力
					for t := 0; t < len(targetStr); t++ {
						ni := i + t*di[dir] + 1
						nj := j + t*dj[dir] + 1
						fmt.Fprintln(w, ni, nj)
					}
					return
				}
			}
		}
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
