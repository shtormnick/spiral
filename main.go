package main

import (
	"fmt"
	//"os"
)

const R int = 4 // capacity of row
const C int = 4 // capacitu of column

func GoldenRatio(matrix [][]int, m, n int) []int { // i - index of first item ,j - index of first column, m - index of last index, n - index of last column
	var  k, l int = 0, 0
	result := make([]int, 0)

	for k < m && l < n {

		for i := l; i < n; i++ {
			result = append(result, matrix[k][i])
		}

		k++

		for i := k; i < m; i++ {
			result = append(result, matrix[i][n-1])
		}

		n--

		if k < m {
			for i := n - 1; i >= l; i-- {
				result = append(result, matrix[m-1][i])
			}
			m--
		}

		if l < n {
			for i := m - 1; i >= k; i-- {
				result = append(result, matrix[i][l])
			}
			l++
		}
	}

	return result
}

func main() {
	input := [][]int{
		[]int{1, 2, 3, 4},
		[]int{4, 5, 6, 7},
		[]int{7, 8, 9, 0},
		[]int{7, 8, 9, 0},
	}
	fmt.Println(GoldenRatio(input, R, C))

}
