package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup //create object for "fork point"
	//goroutines follow fork and join model of concurrency
	wg.Add(1)

	go func() {
		defer wg.Done() // go to "join point"
		data++
	}()

	wg.Wait() // mark "join point"

	if data == 0 {
		fmt.Printf("The value is %v\n", data)
	} else {
		fmt.Println("Race conditions eliminated")
	}
}
