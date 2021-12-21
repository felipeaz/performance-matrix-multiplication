package matrix

const (
	matrixSize = 250
)

func Multiply(matrixA, matrixB [matrixSize][matrixSize]int) [matrixSize][matrixSize]int {
	var result [matrixSize][matrixSize]int
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][row]
			}
		}
	}
	return result
}
