package atomic

import (
	"sync/atomic"
	"testing"
)

// 未调用Store时Load返回nil
func TestValueBase(t *testing.T) {
	var v atomic.Value
	x := v.Load()
	if x == nil {
		t.Logf("atomic.Value init Load value is nil")
	}

	v.Store(64)
	x = v.Load()
	if xx, ok := x.(int); ok {
		t.Logf("atomic.Value Store success , value is %v", xx)
	}
}

//Value 存储的元素类型必须相同，不同的type将导致panic，Store(nil)也会导致panic
func TestValuePanic(t *testing.T) {
	const nilErr = "sync/atomic: store of nil value into Value"
	const badErr = "sync/atomic: store of inconsistently typed value into Value"
	var v atomic.Value
	func() {
		defer func() {
			err := recover()
			if err != nilErr {
				t.Fatalf("inconsistent store panic: got '%v', want '%v'", err, nilErr)
			}
		}()
		v.Store(nil)
	}()
	v.Store(42)
	func() {
		defer func() {
			err := recover()
			if err != badErr {
				t.Fatalf("inconsistent store panic: got '%v', want '%v'", err, badErr)
			}
		}()
		v.Store("foo")
	}()
	func() {
		defer func() {
			err := recover()
			if err != nilErr {
				t.Fatalf("inconsistent store panic: got '%v', want '%v'", err, nilErr)
			}
		}()
		v.Store(nil)
	}()
}
