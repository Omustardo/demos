package solved

import (
	"fmt"
)

// solved on paper
// I forget what the numerical thing is called, but it looks like a triangle:
//     1
//   1   1
//  1  2  1
// 1  3 3  1
//1 4  6  4  1
//
// but then it starts getting smaller
//
// 1 1 1 1 1  1 ... 1 1
// 1 2 3 4 5  6 ... 19 20
// 1 3 6 10 15 21 ...
// 1 4 10 20
// 1 5 15
// .
// .
// 1 20

func Problem15() int {
	var grid [21][21]int // problem is described as a 20x20 grid, but that means it's actually 21x21 points
	
	for i := 0; i < len(grid); i++ {
		grid[0][i] = 1
		grid[i][0] = 1
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j <= i; j++ {
			n := grid[i-1][j] + grid[i][j-1]
			grid[i][j] = n
			grid[j][i] = n
		}
	}
	
	for _, row := range grid {
		fmt.Println(row)
	}
	return grid[20][20]
}
