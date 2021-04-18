package base

import (
	"strconv"
	"testing"
	"time"
)

//针对整个循环的超时
func TestTimeout(t *testing.T) {
	timeout := time.After(time.Second * 3)
	b := make(chan string)
	go SendSomething(b)
	for {
		select {
		case s := <-b:
			t.Logf("case boring chan : %v\n", s)
		case <-timeout:
			t.Logf("timeout return")
			return
		}

	}
}

//针对单个case执行的超时
func TestOneCaseTimeout(t *testing.T) {
	timeout := time.After(time.Second * 3)
	b := make(chan string)
	go SendSomething(b)
	for {
		select {
		case s := <-b:
			t.Logf("case boring chan : %v\n", s)
		case <-timeout:
			t.Logf("timeout return\n")
			return
		case <-time.After(time.Second):
			t.Logf("case timeout return\n")
			return
		}

	}
}

func SendSomething(c chan<- string) {
	for i := 0; i < 20; i++ {
		c <- strconv.Itoa(i)
		time.Sleep(time.Second)
	}
}

func SendSomethingSlow(c chan<- string) {
	for i := 0; i < 20; i++ {
		c <- strconv.Itoa(i)
		time.Sleep(time.Second * 2)
	}
}
