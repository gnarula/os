package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var count = 0

func main() {
	wg.Add(2)
	c := make(chan int)

	go increment("even", c)
	go increment("odd", c)

	c <- count

	wg.Wait()
	fmt.Println("Waited for two threads. Final count value", count, "done")
}

func increment(name string, c chan int) {
	defer wg.Done()
	for {
		select {
		case val, open := <-c:
			if !open {
				return
			}
			if val%2 == 0 {
				if name == "even" {
					count++
					fmt.Println(name, "count =", count)
				}
			} else {
				if name == "odd" {
					count++
					fmt.Println(name, "count =", count)
				}
			}
			if count == 15 {
				close(c)
				return
			}
			c <- count
		}
	}
}
