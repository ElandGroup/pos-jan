// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"
import "time"

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1e10)
		fmt.Println(from, ":", i)
	}
}

func main() {

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	going := "going1"
	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		time.Sleep(1e9)
		fmt.Println(msg)
	}(going)
	fmt.Println("good")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
