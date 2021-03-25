package runtime

import (
	"testing"
)

/*
切片类型源码地址：src/runtime/slice.go

type slice struct {
	array unsafe.Pointer		//指针地址
	len   int					//长度
	cap   int					//容量
}

扩容机制：
容量<1024 - 扩容为原容量的二倍
容量>1024 - 扩容为原容量的1.25倍

使用new创建返回的是一个已清零内存数组指针
使用make创建返回的是结构体，包含数组起点指针，长度和容量
*/

func TestSliceNew(t *testing.T) {
	s := new([]int)
	t.Logf("new [] int , value is %v", s)
	if s == nil {
		t.Error("new [] int , value is nil")
	}

	s2 := make([]int, 100)
	t.Logf("s length is %v", len(s2))
	t.Logf("s length is %v", len(*s))
}
