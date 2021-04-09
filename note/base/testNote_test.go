package base

import (
	"testing"
	"time"
)

func init() {

}

/*
go test -v testNote_test.go
-v 打印具体信息
-cpu=1,2    指定cpu个数，此例子为分别1和2个cpu进行测试。默认使用cpu个数
-count=1    指定执行次数，默认为1
-parallel   设置最大P的个数

对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中只应有一个*testing.T类型的参数声明

结果：
ok      command-line-arguments  10.323s
10.323s   表示   总耗时
*/
func TestFuncNote(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10; i++ {
		t.Logf("banchmark do %v", i)
		time.Sleep(time.Second)
		a := make([]int, 100)
		a = append(a, 1)
	}
}

/*
go test -bench=. -run=^& ./
-bench  表示进行性能测试，值为性能测试函数名称正则， .为任意名称
-run=^$，这个标记用于表明需要执行哪些功能测试函数，这同样也是以函数名称为依据的。
-benchmem 输出基准测试的内存分配统计信息。
-benchtime 用于指定基准测试的探索式测试执行时间上限

该标记的值^$意味着：只执行名称为空的功能测试函数，换句话说，不执行任何功能测试函数。
对于性能测试函数来说，其名称必须以Benchmark为前缀，并且唯一参数的类型必须是*testing.B类型的

结果：
BenchmarkFuncNote-8     1000000000               1.00 ns/op            0 B/op          0 allocs/op
ok      github.com/MSunFlower1014/golang-API/note/base  0.351s
-8 表示 最大P数量为8
1000000000   表示   被测函数的执行次数
1.00 ns/op   表示   被测函数的平均单词执行耗时
B/op 每个操作分配了多少字节。
allocs/op 表示每个操作（单次迭代）发生了多少个不同的内存分配。
0.351s   表示   执行总耗时
*/
func BenchmarkFuncNote(b *testing.B) {
	//忽略 stopTimer 时的耗时
	b.StopTimer()
	for i := 0; i < 1000; i++ {
		b.Logf("banchmark do %v", i)
	}
	b.StartTimer()

	time.Sleep(time.Second)
}

/*
对于示例测试函数来说，其名称必须以Example为前缀，但对函数的参数列表没有强制规定。
*/
func ExampleFuncNote() {

}
