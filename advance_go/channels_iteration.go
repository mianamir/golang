package main

import "fmt"

func integerIterator() <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := 1; i <= 5; i++ {
            ch <- i
        }
    }()
    return ch
}

func main() {
    // Create an integer iterator using a channel.
    iter := integerIterator()

    // Iterate over the values using a for loop.
    for val := range iter {
        fmt.Println(val)
    }
}
