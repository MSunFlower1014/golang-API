package runtime

import (
	"testing"
	"time"
)

/*
chan
1.为nil的管道，发送和接受都会阻塞
2.被关闭的管道，发送会panic，接受会接收到类型的零值
3.关闭nil的管道，会panic
4.关闭已关闭的管道，会panic
5.无缓冲的管道，发送和接受会阻塞到双方完成连接
6.有缓冲的管道，在发送或接受缓冲队列满时进行发送或者接受会阻塞挂起
*/
func TestChan(t *testing.T) {
	m := make(chan int)

	chanSend(m)

	i := chanGet(m)

	if i != 1 {
		t.Fatalf("chan get value is not 1")
	}
}

/*
此参数管道只能用来发送消息，单向发送管道
*/
func chanSend(c chan<- int) {
	c <- 1
}

/*
此参数管道只能用来接受消息，单向接受管道
*/
func chanGet(c <-chan int) int {
	i := <-c
	return i
}

/*
无缓冲管道，发送和接受都会阻塞，直到双方成功建起连接
*/
func TestNoBufferChan(t *testing.T) {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c <- 1
	}()
	before := time.Now()
	i := <-c
	end := time.Now()

	t.Logf("c get i = %v", i)
	t.Logf("c get time is  %v", end.UnixNano()-before.UnixNano())
}

/*
nil管道，发送和接受都会阻塞
*/
func DeadTestNilRecvChan(t *testing.T) {
	c := make(chan int)
	c = nil
	i := <-c
	t.Logf("c get i = %v", i)
}

/*
nil管道，发送和接受都会阻塞
*/
func DeadTestNilSendChan(t *testing.T) {
	c := make(chan int)
	c = nil
	c <- 1
}

/*
有缓冲管道，当发送缓冲满了，后续的发送回阻塞挂起
*/
func TestBufferSendChan(t *testing.T) {
	c := make(chan int, 4)
	for i := 0; i < 4; i++ {
		c <- 1
	}

	go func() {
		time.Sleep(time.Second * 5)
		<-c
	}()
	before := time.Now()
	c <- 1
	end := time.Now()
	t.Logf("c get time is  %v", end.Unix()-before.Unix())
}

/*
有缓冲管道，当接受缓冲为空，会阻塞挂起
*/
func TestBufferRecvChan(t *testing.T) {
	c := make(chan int, 4)
	c <- 1
	before := time.Now()
	i, flag := <-c
	end := time.Now()
	t.Logf("no recv get i = %v , flag = %v ", i, flag)
	t.Logf("c get time is  %v", end.Unix()-before.Unix())
}

/*
已关闭的管道，接受会获得管道类型的零值，flag 为 false
发送会发生panic
*/
func TestCloseChan(t *testing.T) {
	c := make(chan int, 4)
	close(c)
	i, flag := <-c
	t.Logf("no recv get i = %v , flag = %v ", i, flag)

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("cant recover paic ")
		}
		t.Logf("send to closed chan get panic : %v", err)
	}()
	c <- 1
}
