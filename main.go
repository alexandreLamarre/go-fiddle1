package main

import (
	"fmt"
	"time"
)

func somePrint(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Hello world~!")
	//direct call
	somePrint("direct call")

	//goroutine function call
	go somePrint("goroutine-1")

	//goroutine with anonymous function
	go func() {
		somePrint("goroutine-2")
	}()

	//goroutine with function value call
	fv := somePrint
	go fv("goroutine-3")

	// wait for goroutines to end
	fmt.Println("Waiting for goroutines...")
	time.Sleep(time.Millisecond * 100)

	fmt.Println("Done.")
}
