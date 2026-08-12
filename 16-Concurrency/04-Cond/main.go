package main

import (
	"sync"
	"time"
)

var userList = []int{}
var ready = false

func main() {
	condition := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 1000; i++ {
		go NewRequest(i, condition)
	}
	time.Sleep(5 * time.Second)
}
func NewRequest(userId int, condition *sync.Cond) {
	Checking(userId, condition)
	condition.L.Lock()
	defer condition.L.Unlock()
	if !ready {
		condition.Wait()
	}
	println("User ", userId, " start streaming")
}
func Checking(userId int, condition *sync.Cond) {
	println(userId, "waiting for start streaming")
	time.Sleep(time.Millisecond * 150)
	condition.L.Lock()
	defer condition.L.Unlock()
	userList = append(userList, userId)
	if len(userList) == 53 {
		ready = true
		condition.Broadcast()
		println("User ", userId, "waiting end")

	}
}
