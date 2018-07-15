package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readops, writeops uint64
	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("read", atomic.LoadUint64(&readops))
	fmt.Println("write", atomic.LoadUint64(&writeops))
	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()
}
