package sync

import (
	"sync"
	"testing"
	"time"
)

// 需要注意必须保证Signal在Wait后执行，否则可能会空唤醒，导致死锁
func TestCondBase(t *testing.T) {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 3
	done := make(chan int, n)
	running := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func(temp int) {
			c.L.Lock()
			running <- true
			c.Wait()
			t.Logf("go-%v is Signl", temp)
			done <- i
			c.L.Unlock()
		}(i)
	}
	for i := 0; i < n; i++ {
		<-running // Wait for everyone to run.
	}
	m.Lock()
	c.Signal()
	time.Sleep(time.Second * 2)
	c.Broadcast()
	m.Unlock()

	for n > 0 {
		select {
		case r := <-done:
			t.Logf("done num is %v", r)
			n--
		}
	}
}
