package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().String())
		time.Sleep(5 * time.Second)
	}
}
