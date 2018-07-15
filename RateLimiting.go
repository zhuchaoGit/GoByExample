package main

import (
	"time"
	"fmt"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request:", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	requests = make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	for req := range requests {
		<-burstyLimiter
		fmt.Println("req", req, time.Now())
	}
}
