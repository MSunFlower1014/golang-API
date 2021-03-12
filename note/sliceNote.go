package note

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
*/
