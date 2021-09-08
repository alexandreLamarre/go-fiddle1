package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Closure example...")

	var wg sync.WaitGroup
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("Value of i: %v\n", i)
		}()
		fmt.Println("Value of i : ", i)
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("Done.")
}
