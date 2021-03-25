package sync

import (
	"sync"
	"testing"
)

/*
Mutex有两种模式：一般模式和饥饿模式
一般模式:非公平锁，随机获取到锁
饥饿模式：公平锁，根据队列先进先出，且新增得想获取锁的协程不会自旋等待而是直接加入队列等待

当自选超过1ms未获取到锁，会由一般模式转换为饥饿模式
当当前 goroutine 为等待线程得最后一位或者等待时间小于1ms时，由饥饿模式转换为一般模式
*/
func TestMutexBase(t *testing.T) {
	m := new(sync.Mutex)
	/*
		1.使用CompareAndSwapInt32更新state值
		2.更新成功表示获取锁成功
		3.失败则开始 lockSlow
		4.lockSlow 为循环尝试获取锁
	*/
	m.Lock()
	func() {
		t.Logf("do something")
	}()
	m.Unlock()
}
