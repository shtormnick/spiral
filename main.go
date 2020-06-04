package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func GoldenRatio (matrix [][]int) []int {
	size := len(matrix)
	start_point := Point{x : int(size/2), y: int(size/2) }
	end_point := Point{x: 1, y: size }
	direction := 0 												//current direction; 0=RIGHT, 1=DOWN, 2=LEFT, 3=UP
	counter := 0
	chain_size := 1
	x := 0
	y := 0
	fmt.Println("start point: ", start_point )
	fmt.Println("end point: ", end_point )
	for k := 1; k <= ( size - 1 ); k++ {
		if k < (size - 1){
			for j := 0; j < 2; j++ {
				for i := 0; i < chain_size; i++ {
					fmt.Printf("%v ", matrix[x][y])
					counter++
					switch direction {
						case 0: y = y + 1
						case 1: x = x + 1
						case 2: y = y - 1
						case 3: x = x - 1
					}
				}
				direction = (direction+1)%4;
			}
		} else {
			for j := 0; j < 3; j++ {
				for i := 0; i < chain_size; i++ {
					fmt.Printf("%v ", matrix[x][y])
					counter++
					switch direction {
						case 0: y = y + 1
						case 1: x = x + 1
						case 2: y = y - 1
						case 3: x = x - 1
					}
				}
				direction = (direction+1)%4;
			}
		}
        chain_size = chain_size + 1;
    }
	return make([]int, 0)
}


func main() {
	input := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	fmt.Println(GoldenRatio(input))
}