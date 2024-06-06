// Section 2: Concurrency

// Channels
package section2

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(ch chan int) {
    defer close(ch)
    count := 0
    for count < 5 {
        num := rand.Intn(100)
        ch <- num
        count++
        time.Sleep(time.Second)
    }
}

func Consumer(ch chan int) {
	for num := range ch {
		fmt.Printf("Received %d, its square is %d\n", num, num*num)
	}
}
