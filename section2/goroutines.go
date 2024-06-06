// Section 2: Concurrency

// Goroutines
package section2

import (
	"sync"
)

func SumArray(arr []int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	sum := 0
	for _, v := range arr {
		sum += v
	}
	resultChan <- sum
}
