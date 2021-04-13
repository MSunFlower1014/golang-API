package sync

import (
	"sync"
	"testing"
)

/*
sync.Map适合读多写少得场景
map的key不能是函数类型，字段类型和切片类型
map结构体中包含两个原生得字典 - readOnly 和 dirty
1. 空间换时间。 通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响
2. 使用只读数据(read)，避免读写冲突。
3. 动态调整，miss次数多了之后，将dirty数据提升为read
4. 延迟删除。 删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据
5. 优先从read读取、更新、删除，因为对read的读取不需要锁

amended为true时表示一些key存在于dirty但是不存在于read中，代表read部分过时

当 dirty 为 nil 的时候，read 就代表 map 所有的数据；当 dirty 不为 nil 的时候，dirty 才代表 map 所有的数据。
*/
func TestSyncMapBase(t *testing.T) {
	k := "1"
	v := "v1"
	m := new(sync.Map)
	/*
		1.查询 read 字典中是否存在，存在则常识替换value值并返回
		2.如果 read 中不存在或者替换value失败，则加锁
		3.重新判断 read 中是否存在
		4.如果存在 read 中且倍标记为已删除，则直接存储在 dirty 中
		5.如果不存在 read 中而存在 dirty 中，也更新存储到 dirty 中
		6.如果在 read 和 dirty 中都不存在，则新建 entry 存入 dirty并修改amended为true标记dirty中有部分key在read中不存在
		7.解锁
	*/
	m.Store(k, v)

	/*
		1.查询 read 中是否存在key
		2.不存在于 read 中时加锁后重新尝试
		3.仍然不存在 read 且amended为true时查询 dirty ，当在dirty中获取成功，如果miss次数++，当达到一定限度会 read-dirty转换
		4.如果都获取失败，返回 nil 和 false
		5.获取成功则返回对应 value
		read-dirty转换：当多此查询在read 中未找到，而在 dirty中获取到，则 dirty升级为 read ，并将新dirty置为nil
	*/
	value, ok := m.Load(k)
	if !ok {
		t.Errorf("load k error , key is %v", k)
	}
	t.Logf("load k : %v , value : %v", k, value)

	value, ok = m.Load("not ok")
	if !ok {
		t.Logf("error key result value is %v", value)
	}

	k = "2"
	v = "v2"
	value, ok = m.LoadOrStore(k, v)
	//ok表示是否loaded成功，当map中存在该key时返回value且loaded为true
	t.Logf("LoadOrStore result value : %v , loaded is %v", value, ok)
	value, ok = m.LoadOrStore(k, v)
	t.Logf("Second LoadOrStore result value : %v , loaded is %v", value, ok)

	m.Delete(k)
	value, ok = m.LoadOrStore(k, v)
	t.Logf("Third LoadOrStore result value : %v , loaded is %v", value, ok)
	//range时如果发现amended为true也会进行read-dirty转换
	m.Range(func(key, value interface{}) bool {
		t.Logf("range key is %v value is %v", key, value)
		return true
	})
}

const (

	// flags
	iterator     = 1 // there may be an iterator using buckets
	oldIterator  = 2 // there may be an iterator using oldbuckets
	hashWriting  = 4 // a goroutine is writing to the map
	sameSizeGrow = 8 // the current map growth is to a new map of the same size

)

func TestUnitBase(t *testing.T) {
	var flags uint
	if flags&hashWriting != 0 {
		t.Errorf("concurrent map writes")
	}
	flags ^= hashWriting
	t.Logf("flags ^= hashWriting value is %v", flags)
	if flags&hashWriting == 0 {
		t.Errorf("concurrent map writes")
	}
	flags &^= hashWriting
	t.Logf("flags &^= hashWriting value is %v", flags)
}
