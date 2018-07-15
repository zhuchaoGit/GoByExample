package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var ops uint64
	var comp int
	for i := 0; i < 50; i++ {
		go func() {
			for {
				comp++
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("ops", atomic.LoadUint64(&ops), "cmp", comp)
}
