package main

import (
	"fmt"
	"matrix-multiplication/matrix"
	"time"
)

func main() {
	fmt.Println("Started multiplication")

	start := time.Now()
	matrixA := matrix.GenerateSquareMatrix()
	matrixB := matrix.GenerateSquareMatrix()
	matrix.Multiply(matrixA, matrixB)

	elapsed := time.Since(start)
	fmt.Printf("Process took %s\n", elapsed)
}
