package sync

import (
	"sync"
	"testing"
)

func TestOnceBase(t *testing.T) {
	o := new(sync.Once)
	n := 10
	done := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			o.Do(func() {
				t.Logf("once done")
			})
			done <- true
		}()
	}

	for n > 0 {
		select {
		case <-done:
			n--
		}
	}
}
