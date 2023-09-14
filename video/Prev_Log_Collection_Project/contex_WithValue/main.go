package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"time"
)

type TraceCode string
type UserID string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	useridKey := UserID("USER_ID")
	userid, ok := ctx.Value(useridKey).(int64)
	if !ok {
		fmt.Println("invalid userid code")
	}
	log.Printf("%s worker func", traceCode)
	log.Printf("userid: %d", userid)
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		fmt.Printf("worker, userid:%s\n", userid)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	ctx = context.WithValue(ctx, UserID("USER_ID"), int64(21341313213213132))
	log.Printf("%s main func", "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
