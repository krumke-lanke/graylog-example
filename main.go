package main

import (
	"time"
	"fmt"
)

func main() {

	i := 0
	for range time.Tick(5 * time.Second) {
		i++
		fmt.Println("hello world ", i)
	}


}
