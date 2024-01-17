package main

import (
	"fmt"
	"time"
)

func printMessage(msg string, channel chan string) {
	for i := 0; i < 5; i++ {
		//fmt.Println("Sleeping ", msg)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Printing ", i, msg)
	}
	channel <- "done"
}
func main() {
	channel := make(chan string)
	go printMessage("Thread1", channel)
	go printMessage("T2", channel)
	fmt.Println(<-channel)
	//fmt.Println(resp)
	//printMessage("Thread2")
	// printMessage("Thread3")
}
