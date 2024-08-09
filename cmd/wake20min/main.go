package main

import (
	"fmt"
	"time"

	"github.com/pjsoftware/go-wake"
)

func main() {
	fmt.Printf("Setting Wakefulness!\n")
	wake.KeepScreenOn()

	min := 0
	for {
		fmt.Printf("Awake at minute %d\n", min)
		time.Sleep(60 * time.Second)
		min++

		if min == 15 {
			fmt.Printf("Allowing sleep!\n")
			wake.AllowSleep()
		}
	}
}