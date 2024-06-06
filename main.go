package main

import (
	"fmt"
	"sync"
	"time"
	"os"
	"go-test-exercise/section1"
	"go-test-exercise/section2"
	"go-test-exercise/section3"
)

func main() {
	// Section 1: Variables and Types
	x, y := 1, 2
	x, y = section1.Swap(x, y)
	fmt.Println("Swapped values:", x, y) // Output: 2 1

	// Section 1: Slices
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Sum of even numbers:", section1.SumEven(numbers)) // Output: 12

	// Section 1: Interfaces
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fileLogger := &section1.FileLogger{File: file}
	consoleLogger := &section1.ConsoleLogger{}

	fileLogger.Log("This is a file log message.")
	consoleLogger.Log("This is a console log message.")

	// Section 2: Goroutines
	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr[i] = i + 1
	}

	numGoroutines := 4
	partSize := len(arr) / numGoroutines
	resultChan := make(chan int, numGoroutines)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		wg.Add(1)
		go section2.SumArray(arr[start:end], &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Println("Total sum:", totalSum)

	// Section 2: Channels
	ch := make(chan int)
	go section2.Producer(ch)
	go section2.Consumer(ch)

	time.Sleep(10 * time.Second)

	// Section 3: HTTP Server
	section3.StartServer()
	time.Sleep(1 * time.Second)
}

