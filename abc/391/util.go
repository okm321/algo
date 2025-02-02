package _391

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

// readInt 単一の整数を読み取る関数
func readInt(reader *bufio.Reader) int {
	line, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(line))
	return val
}

func readBigInt(reader *bufio.Reader) *big.Int {
	line, _ := reader.ReadString('\n')
	val, _ := strconv.Atoi(strings.TrimSpace(line))
	return big.NewInt(int64(val))
}

// readString 単一の文字列を読み取る関数
func readString(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
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

func readAnyArray(reader *bufio.Reader) []any {
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	arr := make([]any, len(parts))
	for i, s := range parts {
		arr[i] = strings.TrimSpace(s)
	}

	return arr
}

// readStringArray 文字列の配列を読み取る関数
func readStringArray(reader *bufio.Reader) []string {
	line, _ := reader.ReadString('\n')
	return strings.Fields(line)
}

// height行の文字列グリッドを読み込む
func readGrid(r *bufio.Reader, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		str := readString(r)
		grid[i] = strings.Split(str, "")
	}
	return grid
}

// 一行に2つの整数のみの入力を読み込む
func read2Ints(r *bufio.Reader) (int, int) {
	input, _ := r.ReadString('\n')
	strs := strings.Fields(input)
	i1, _ := strconv.Atoi(strs[0])
	i2, _ := strconv.Atoi(strs[1])
	return i1, i2
}

// height行の整数グリッドを読み込む
func readIntGrid(r *bufio.Reader, height int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = readIntArray(r)
	}

	return grid
}
