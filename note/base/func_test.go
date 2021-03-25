package base

import "testing"

func TestFunc(t *testing.T) {
	n := 1000
	done := make(chan int, n)
	for i := 0; i < n; i++ {
		j := i
		go func() {
			t.Logf(" i value is %v", j)
			done <- j
		}()
	}
	m := make(map[int]int)
	for n > 0 {
		select {
		case c := <-done:
			n--
			m[c]++
		}
	}
}

//不创建循环内局部变量时i会逃逸到堆中，go func 调用的i为同一个
func TestFunc2(t *testing.T) {
	n := 1000
	done := make(chan int, n)
	for i := 0; i < n; i++ {
		go func() {
			t.Logf(" i value is %v", i)
			done <- i
		}()
	}
	m := make(map[int]int)
	for n > 0 {
		select {
		case c := <-done:
			n--
			m[c]++
		}
	}
}
