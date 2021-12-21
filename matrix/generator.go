package matrix

import (
	"math/rand"
	"time"
)

func GenerateSquareMatrix() [matrixSize][matrixSize]int {
	var matrix [matrixSize][matrixSize]int
	rand.Seed(time.Now().UnixNano())
	for line := 0; line < matrixSize; line++ {
		for column := 0; column < matrixSize; column++ {
			matrix[line][column] += rand.Intn(20) - 20
		}
	}

	return matrix
}
