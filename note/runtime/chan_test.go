package runtime

import "testing"

/*

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
