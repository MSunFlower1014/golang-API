package atomic

import (
	"sync/atomic"
	"testing"
)

func TestInt64Swap(t *testing.T) {
	var a int64 = 0
	var i *int64 = &a
	atomic.StoreInt64(i, a)
	old := atomic.SwapInt64(i, 2)
	t.Logf(" * int64 init value a is %v", old)
	if 2 == *i {
		t.Logf("SwapInt64 is success")
	}
	flag := atomic.CompareAndSwapInt64(i, 2, 2)
	t.Logf("CompareAndSwapInt64 result is %v", flag)
	old = atomic.AddInt64(i, 3)
	t.Logf("after AddInt64 value is %v", *i)

	newI := atomic.LoadInt64(i)
	t.Logf("LoadInt64 return value is %v", newI)

	var b int64 = 3
	bAddResult := atomic.AddInt64(&b, 3)
	t.Logf("bAddResult value is %v", bAddResult)
}
