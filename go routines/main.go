package main

import (
	"fmt"
	"time"
)

// Race conditions

func main() {
	fmt.Println("starting main")
	msg := "hi"
	go func() {
		loop(msg)
	}()
	msg = "yo"
	fmt.Println("End main")
	time.Sleep(100 * time.Millisecond)

}

func loop(m string) {
	for i := 0; i < 5; i++ {
		fmt.Println(m)
	}
}
