package matrix

import (
	"fmt"
	"sync"
	"time"
)

const (
	size = 250
)

var (
	matrixA [size][size]int
	matrixB [size][size]int
	result  [size][size]int

	rwLocker  = sync.RWMutex{}
	cond      = sync.NewCond(rwLocker.RLocker())
	waitGroup = sync.WaitGroup{}
)

func Start() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		GenerateSquareMatrix(&matrixA)
		GenerateSquareMatrix(&matrixB)
		for row := 0; row < size; row++ {
			Multiply(row)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Process took %s\n", elapsed)
}

func Multiply(row int) {
	for col := 0; col < size; col++ {
		for i := 0; i < size; i++ {
			result[row][col] += matrixA[row][i] * matrixB[i][row]
		}
	}
}

func StartProcess() {
	// Throws one thread for each matrix row
	waitGroup.Add(size)
	for row := 0; row < size; row++ {
		go PerformanceMultiply(row)
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		waitGroup.Wait()
		rwLocker.Lock()
		GenerateSquareMatrix(&matrixA)
		GenerateSquareMatrix(&matrixB)
		waitGroup.Add(size)
		rwLocker.Unlock()
		cond.Broadcast()
	}
	elapsed := time.Since(start)
	fmt.Printf("Process took %s\n", elapsed)
}

func PerformanceMultiply(row int) {
	rwLocker.RLock()
	for {
		waitGroup.Done()
		cond.Wait()
		for col := 0; col < size; col++ {
			for i := 0; i < size; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][row]
			}
		}
	}
}
