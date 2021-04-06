package base

import (
	"runtime"
	"testing"
	"time"
)

/*
设置环境变量开启 gc日志
GOGCTRACE=1;GODEBUG=gctrace=1

linux中
GODEBUG=gctrace=1 执行文件即可

格式

gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
含义

	gc #        GC次数的编号，每次GC时递增
	@#s         距离程序开始执行时的时间
	#%          GC占用的执行时间百分比
	#+...+#     GC使用的时间
	#->#-># MB  GC开始，结束，以及当前活跃堆内存的大小，单位M
	# MB goal   全局堆内存大小
	# P         使用processor的数量
如果每条信息最后，以(forced)结尾，那么该信息是由runtime.GC()调用触发
*/
func TestTc(t *testing.T) {
	var s []int
	for i := 0; i < 1000; i++ {
		s = append(s, i)
	}

	t.Logf("force gc")
	runtime.GC()
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	t.Logf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
	time.Sleep(time.Second * 3600)
}
