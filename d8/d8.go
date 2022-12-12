package d8

import (
	"fmt"

	. "github.com/mfigurski80/AOC22/utils"
)

func findVisibility(row []uint8) []bool {
	viz := make([]bool, len(row))
	left := 0
	left_max := row[left]
	right := len(row) - 1
	right_max := row[right]
	for left <= right {
		for ; left_max <= right_max && left <= right; left++ {
			if row[left] > left_max {
				left_max = row[left]
				viz[left] = true
				continue
			}
			viz[left] = false
		}
		for ; left_max >= right_max && left <= right; right-- {
			if row[right] > right_max {
				right_max = row[right]
				viz[right] = true
				continue
			}
			viz[right] = false
		}
	}

	viz[0] = true
	viz[len(row)-1] = true
	return viz
}

func findHeightMapVisibility(heightMatrix [][]uint8) [][]bool {
	vizMatrix := make([][]bool, len(heightMatrix))
	// find visibility from sides
	for i, row := range heightMatrix {
		vizMatrix[i] = findVisibility(row)
	}
	// find visibility from top/bottom
	for i := 1; i < len(heightMatrix)-1; i++ {
		col := make([]uint8, len(heightMatrix))
		for j, row := range heightMatrix {
			col[j] = row[i]
		}
		viz := findVisibility(col)
		for j, row := range vizMatrix {
			row[i] = row[i] || viz[j]
		}
	}
	return vizMatrix
}

func findScenicScore(row []uint8) []uint {
	scenic := make([]uint, len(row))
	for i, height := range row {
		// check left
		left_scenic := uint(0)
		for j := i - 1; j >= 0; j-- {
			left_scenic++
			if row[j] >= height {
				break
			}
		}
		// check right
		right_scenic := uint(0)
		for j := i + 1; j < len(row); j++ {
			right_scenic++
			if row[j] >= height {
				break
			}
		}
		scenic[i] = left_scenic * right_scenic
	}
	return scenic
}

func findHeightMapScenicScore(heightMatrix [][]uint8) [][]uint {
	scenicMatrix := make([][]uint, len(heightMatrix))
	// find scenic from sides
	for i, row := range heightMatrix {
		scenicMatrix[i] = findScenicScore(row)
	}
	// find scenic from top/bottom
	for i := range heightMatrix {
		col := make([]uint8, len(heightMatrix))
		for j, row := range heightMatrix {
			col[j] = row[i]
		}
		scenic := findScenicScore(col)
		for j, val := range scenic {
			scenicMatrix[j][i] *= val
		}
	}
	return scenicMatrix
}

func Main() {
	// make 2d uint matrix
	heightMatrix := make([][]uint8, 0)
	width := 0

	// read file into matrix
	DoByFileLine("d8/in.txt", func(line string) {
		if width == 0 {
			width = len(line)
		}
		row := make([]uint8, width)
		for i, char := range line {
			row[i] = uint8(char - '0')
		}
		heightMatrix = append(heightMatrix, row)
	})

	// vizMatrix := findHeightMapVisibility(heightMatrix)
	// // count visible trees
	// count := 0
	// for _, row := range vizMatrix {
	// 	for _, visible := range row {
	// 		if visible {
	// 			count++
	// 		}
	// 	}
	// }
	// fmt.Println(count)

	// fmt.Println(heightMatrix)
	scenicMatrix := findHeightMapScenicScore(heightMatrix)
	// fmt.Println(scenicMatrix)
	// find max scenic score
	max := uint(0)
	for _, row := range scenicMatrix {
		for _, score := range row {
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)

	// fmt.Println(findScenicScore([]uint8{1, 0, 0, 2, 2, 1, 2, 0, 0, 0, 3, 2, 0, 2, 3, 2, 1, 0, 2, 3}))
}
