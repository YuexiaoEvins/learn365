package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			select {
			case intCh <- 1:
			case intCh <- 2:
			case intCh <- 3:
			}
		}
		fmt.Println("end. [sender]")
	}()

	count := 0
	for {
		if elem, ok := <-intCh; ok {
			count += elem
			if count > 10 {
				break
			}
		} else {
			break
		}
	}

	fmt.Println("receiver quit with count:", count)

}
