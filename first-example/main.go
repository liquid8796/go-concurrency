package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(9)

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x))
	}

	printSomething("This is the second thing to be printed!")
}
