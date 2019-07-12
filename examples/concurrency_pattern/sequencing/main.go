package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boring(msg string) <-chan Message { // Returns receive-only channel of strings.
	c := make(chan Message)
	waitForIt := make(chan bool) // Give main control over our execution.
	go func() {                  // Launch the goroutine from inside the function. Function Literal.
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt // Block until main tells us to go again.
		}
	}()
	return c // Return the channel to the caller.
}
func fanIn(input1, input2 <-chan Message) chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
func main() {
	rand.Seed(time.Now().UnixNano())
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

}
