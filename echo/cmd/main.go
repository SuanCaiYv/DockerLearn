package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello, Golang!")
		time.Sleep(5 * time.Second)
	}
}
