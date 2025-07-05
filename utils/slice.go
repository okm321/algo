package utils

// Compact 同じ値が連続する部分を1つにまとめる関数
func Compact(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	result := make([]int, 0, len(arr))
	result = append(result, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}

	return result
}

// 標準ライブラリのみを使った要素の検索
func IndexOf(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
