package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func receive(strChan <-chan string,
	syncChan <-chan struct{},
	syncChan2 chan<- struct{},
) {
	<-syncChan
	fmt.Println("receive a sync signal and wait a second... [receiver]")

	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string,
	syncChan chan<- struct{},
	syncChan2 chan<- struct{},
) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan <- struct{}{}
			fmt.Println("Send a sync Signal. [sender]")

		}
	}

	fmt.Println("wait 2 seconds... [sender]")
	time.Sleep(2 * time.Second)
	close(strChan)
	syncChan2 <- struct{}{}

}

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan, syncChan1, syncChan2)
	go send(strChan, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2
}
