package runtime

import (
	"fmt"
	"github.com/MSunFlower1014/golang-API/pkg/model"
	"reflect"
	"testing"
)

/*
接口由两个元素组成：类型和值
var v *model.Book
var i interface{}
i = v
此时 i 为 (*model.Book)(nil) , 类型为 model.Book 的指针类型，值为nil
所以 i == nil is false

函数参数为error接口，所以此时只有值为nil，为类型不为nil

可以使用
fmt.Printf("b.value == nil is %t\n", b == nil || (reflect.ValueOf(b).Kind() == reflect.Ptr && reflect.ValueOf(b).IsNil()))
判断接口是否为nil

优化nil判断参考  isNilFixed 方法
*/
type MyError struct {
	errCode uint8
}

func (e *MyError) Error() string {
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}

func TestCheckError(t *testing.T) {
	var e *MyError
	checkError(e)
}

func TestInterfaceNil(t *testing.T) {
	var v *model.Book
	var i interface{}
	i = v
	if i == nil {
		t.Error(" interface is nil ")
	}
	t.Logf("interface not nil , value : %v", i)
}

func checkError(e error) {
	fmt.Printf("b.value == nil is %t\n", e == nil || (reflect.ValueOf(e).Kind() == reflect.Ptr && reflect.ValueOf(e).IsNil()))
	if e != nil {
		panic(e)
	}
}

//优化的nil判断
func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
