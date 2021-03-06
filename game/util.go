package game

import (
	"time"
	"math/rand"
	"log"
	"os"
)

// Compress align the non blank field to a side.
func Compress(field [][]int) [][]int {
	ret := make([][]int, len(field))
	for i, v := range field {
		row := make([]int, len(v))
		idx := 0
		for _, v := range v {
			if v == 0 {
				continue
			}
			row[idx] = v
			idx++
		}
		ret[i] = row
	}
	return ret
}

// Invert swap the value of all rows in field from left to right.
func Invert(field [][]int) [][]int {
	rows := len(field)
	if rows == 0 {
		return field

	}
	ret := make([][]int, rows)
	cols := len(field[0])
	for r, row := range field {
		newRow := make([]int, cols)
		for c, v := range row {
			newRow[cols-c-1] = v
		}
		ret[r] = newRow
	}
	return ret
}

// Transpose the two dimension slice just like the transpose operator of matrix.
func Transpose(field [][]int) [][]int {
	rows := len(field)
	cols := len(field[0])
	ret := make([][]int, cols)
	for i := 0; i < cols; i++ {
		row := make([]int, rows)
		for j := 0; j < rows; j ++ {
			row[j] = field[j][i]
		}
		ret[i] = row
	}
	return ret
}

// Random returns the number between min and max
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

var logger *log.Logger

func init() {
	file, _ := os.OpenFile("dev.log", os.O_RDWR|os.O_CREATE, 0600)
	logger = log.New(file, "", log.Llongfile|log.LstdFlags)
}