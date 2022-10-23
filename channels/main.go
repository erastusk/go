package main

/* Channels provide a communication channel for go routines, where other routines can read and write into the channel*/

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println("Main starts")
	// ch channel receives value 2
	ch <- 2

	/* ch channel needs to send that value to another go routine, in our case main
	Why deadlock errors?
	All channel receiving operations are "blocking" in nature, so the go routine runing the main function got blocked at line 12. Main
	go routine writes into channel, hence at that moment it gets blocked and another routine is attempted to be spun up.
	When blocked, the go scheduler tries to schedule another go routine but in our case there's none, resulting into a
	sleep - deadlock error.
	*/
	//<-ch
	fmt.Println("Main ends")
	time.Sleep(50 * time.Second)
	<-ch

}
