package main

import "fmt"

func main() {
	var data int

	//Race condition occurs when order of execution is NOT guaranteed

	//Concurrent programs does not execute in the order they are coded
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("The value is %v\n", data)
	}

}
