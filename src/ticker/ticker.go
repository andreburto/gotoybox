package main

import (
	"fmt"
	"time"
)

func HelloTicker() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Hello, ticker!")
	}
}

func WorldTicker() {
	ticker := time.NewTicker(7 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("World, ticker!")
	}
}

func main() {
	go HelloTicker()
	go WorldTicker()

	select {}
}
