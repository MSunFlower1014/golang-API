package sync

import (
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
)

/*
fmt包通过pool提高效率
fmt/print.go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func newPrinter() *pp {
	//从pool中获取pp
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.wrapErrs = false
	p.fmt.init(&p.buf)
	return p
}

func (p *pp) free() {
	// 正确使用sync.Pool要求每个条目具有大约相同的内存成本。
	//为了在存储的类型包含可变大小的缓冲区时获得此属性，
	//我们对最大缓冲区添加了硬限制以放回池中。

	// See https://golang.org/issue/23199
	if cap(p.buf) > 64<<10 {
		return
	}

	p.buf = p.buf[:0]
	p.arg = nil
	p.value = reflect.Value{}
	p.wrappedErr = nil
	//清空后放回pool
	ppFree.Put(p)
}
*/
func TestPoolBase(t *testing.T) {
	// disable GC so we can control when it happens.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var p sync.Pool
	if p.Get() != nil {
		t.Fatal("expected empty")
	}

	p.Put("a")
	p.Put("b")
	if g := p.Get(); g != "a" {
		t.Fatalf("got %#v; want a", g)
	}
	if g := p.Get(); g != "b" {
		t.Fatalf("got %#v; want b", g)
	}
	if g := p.Get(); g != nil {
		t.Fatalf("got %#v; want nil", g)
	}

	// Put in a large number of objects so they spill into
	// stealable space.
	for i := 0; i < 100; i++ {
		p.Put("c")
	}
	// After one GC, the victim cache should keep them alive.
	runtime.GC()
	if g := p.Get(); g != "c" {
		t.Fatalf("got %#v; want c after GC", g)
	}
	// A second GC should drop the victim cache.
	runtime.GC()
	if g := p.Get(); g != nil {
		t.Fatalf("got %#v; want nil after second GC", g)
	}
}
