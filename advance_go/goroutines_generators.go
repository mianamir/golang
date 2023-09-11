package main

import (
	"fmt"
	"time"
)

func fibonacciGenerator() <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        a, b := 0, 1
        for {
            ch <- a
            a, b = b, a+b
            time.Sleep(time.Millisecond * 500) // Simulate work
        }
    }()
    return ch
}

func main() {
    // Create a Fibonacci number generator using a channel.
    gen := fibonacciGenerator()

    // Consume values from the generator.
    for i := 0; i < 10; i++ {
        fmt.Println(<-gen)
    }
}
