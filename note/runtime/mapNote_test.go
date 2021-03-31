package runtime

import "testing"

/*
map 类型源码地址：src/runtime/map.go

type hmap struct {
	count     int // 数据总量
	B         uint8  // 容量为2的B次方
	hash0     uint32 // hash种子

	buckets    unsafe.Pointer // 程度为2的B次方的桶数组，当count为0时可能为nil
	oldbuckets unsafe.Pointer // 扩容时旧数据保存桶数组，仅仅当扩容时不为nil
}

make map 流程：
new(hmap) - 计算并设置hash0和B - 创建桶数组

扩容:新建2倍容量的桶数组，并包留旧桶数组
哈希表数据的实际复制是以增量方式完成的
通过growWork（）和excavate（）

get:根据高位hash获取桶数组下标，低八位hash通过比较后获取key-value（每个桶内8个元素，一次比较速度也很快）
先查询旧桶数组，后查询新桶数组
*/

func TestMap(t *testing.T) {
	m := make(map[string]string, 16)
	t.Logf("map length is %v ", len(m))
	if m == nil {
		t.Error("")
	}
	m["1"] = ""
	t.Logf("map length is %v ", len(m))
	//t.Logf("map cap is %v ", cap(m))
}
