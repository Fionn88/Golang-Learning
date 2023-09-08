package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker2(ctx context.Context) {
LABEL:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LABEL
		default:
		}
	}
}

func worker(ctx context.Context) {
	go worker2(ctx)
	defer wg.Done()
LABEL:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:
		}
	}

}

func main() {
	cxt, canncel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(cxt)

	time.Sleep(time.Second * 5)
	canncel()
	wg.Wait()
	fmt.Println("canncel...")
	time.Sleep(time.Second * 5)
	fmt.Println("over")
}
