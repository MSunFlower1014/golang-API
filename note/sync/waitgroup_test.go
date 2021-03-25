package sync

import (
	"sync"
	"testing"
	"time"
)

func TestWaiteGroupBase(t *testing.T) {
	w := new(sync.WaitGroup)
	n := 10
	w.Add(n)

	for i := 0; i < n; i++ {
		temp := i
		go func() {
			time.Sleep(time.Second * 1)
			t.Logf("go func num is %v", temp)
			w.Done()
		}()
	}
	w.Wait()
}
