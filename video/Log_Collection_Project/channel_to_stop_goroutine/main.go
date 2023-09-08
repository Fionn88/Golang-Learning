package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ch <-chan bool) {
	defer wg.Done()
LABEL:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ch:
			break LABEL
		default:
		}
	}

}

func main() {
	var exitChan = make(chan bool)
	wg.Add(1)
	go worker(exitChan)

	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
	fmt.Println("over")
}
