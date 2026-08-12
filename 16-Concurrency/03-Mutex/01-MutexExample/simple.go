package main

import (
	"sync"
)

func SimpleGood() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	Counter := 0
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			mx.Lock()
			Counter++
			mx.Unlock()
		}()
	}
	wg.Wait()
	println(Counter)
}
func SimpleBad() {
	wg := sync.WaitGroup{}
	Counter := 0
	wg.Add(1000000)
	for i := 0; i < 1_000_000; i++ {
		go func() {
			defer wg.Done()
			Counter++
		}()
	}
	wg.Wait()
	println(Counter)
}
