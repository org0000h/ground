package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1.2e3)) * time.Millisecond)
				// do nothing
			case <-quit:
				// cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}
