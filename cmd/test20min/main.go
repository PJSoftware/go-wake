package main

import (
	"fmt"
	"time"
)

func main() {
	min := 0
	for {
		fmt.Printf("Awake at minute %d\n", min)
		time.Sleep(60 * time.Second)
		min++
	}
}
