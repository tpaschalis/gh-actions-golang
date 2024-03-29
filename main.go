package main

import "fmt"
import "time"
import "github.com/Pallinder/go-randomdata"

func main() {
	fmt.Println("Running for GH Actions!")
	fmt.Println("A round of applause for", randomdata.SillyName())

	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
	fmt.Println(fibonacci(8))
}

func fibonacci(n int) int {
	time.Sleep(10 * time.Millisecond) // Evil code
	cur := 0
	i, j := 0, 1
	for cur < n {
		i, j = j, i+j
		cur++
	}
	return j
}
