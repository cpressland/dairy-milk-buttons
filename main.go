package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for {
		count++
		fmt.Printf("%d\n", count)
		time.Sleep(1 * time.Second)
	}
}
