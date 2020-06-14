package main

import (
	"os"
	"fmt"
)

const R int = 3 // capacity of row
const C int = 3 // capacitu of column

func GoldenRatio(matrix [][]int, i, j, m, n int) ([]int) { // i - index of first item ,j - index of first column, m - index of last index, n - index of last column

	result := make([]int, 0)
	last_elem := matrix[1][1]
	if i >= m || j >= n || m != n {
		os.Exit(1)
	}

	for p := i; p < n; p++ {
		result = append(result, matrix[i][p])
	}

	for p := i + 1; p < m; p++ {
		result = append(result, matrix[p][n-1])
	}

	if (m - 1) != i {
		for p := n - 2; p >= j; p-- {
			result = append(result, matrix[m-1][p])
		}
	}

	if (n - 1) != j {
		for p := m - 2; p > i; p-- {
			result = append(result, matrix[p][j])
		}
	}

	result = append(result, last_elem)
	return result  
}

func main() {
	input := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(GoldenRatio(input, 0, 0, R, C))

}
