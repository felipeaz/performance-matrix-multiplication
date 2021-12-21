package matrix

import (
	"math/rand"
	"time"
)

func GenerateSquareMatrix(matrix *[size][size]int) {
	rand.Seed(time.Now().UnixNano())
	for line := 0; line < size; line++ {
		for column := 0; column < size; column++ {
			matrix[line][column] += rand.Intn(20) - 20
		}
	}
}
