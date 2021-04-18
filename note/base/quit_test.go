package base

import (
	"fmt"
	"testing"
)

func TestQuit(t *testing.T) {

	quit := make(chan string)
	send := make(chan int)
	go Doing(quit, send)

	for i := 0; i < 10; i++ {
		send <- i
	}
	//发送消息告知执行结束
	quit <- "execute end"
	//等待释放资源结束
	t.Logf("quit get :%v\n", <-quit)

}

func Doing(quit chan string, iC chan int) {
	for {
		select {
		case s := <-quit:
			fmt.Printf("string chan get  : %v \n", s)
			//释放资源
			quit <- "bye"
		case i := <-iC:
			fmt.Printf("int chan get : %v\n", i)

		}
	}
}
