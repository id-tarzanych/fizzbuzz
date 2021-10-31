package main

import "fmt"

const LIMIT = 25

func fizzBuzz(n int, ch chan interface{}) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			ch <- "Fizz Buzz"
		case i%3 == 0:
			ch <- "Fizz"
		case i%5 == 0:
			ch <- "Buzz"
		default:
			ch <- i
		}
	}

	close(ch)
}

func main() {
	ch := make(chan interface{})

	go fizzBuzz(LIMIT, ch)

	for v := range ch {
		fmt.Println(v)
	}
}
