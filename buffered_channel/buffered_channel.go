package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// will be execute

	ch <- 3
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// will be execute

	ch <- 5
	ch <- 6
	ch <- 7
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// will error deadlock

}
