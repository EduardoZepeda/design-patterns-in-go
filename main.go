package main

import (
	"fmt"
	"sync"
)

func printHello(wg *sync.WaitGroup) {
	fmt.Println("H")
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printHello(&wg)
	}
	wg.Wait()
	fmt.Println("Hello World")
}
