package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
	time.Sleep(1 * time.Second)
}

func main() {
	go printSomething("This is the first thing to be printed!")

	printSomething("This is the second thing to be printed!")
}
