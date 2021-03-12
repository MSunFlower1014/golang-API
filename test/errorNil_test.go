package test

import (
	"fmt"
	"testing"
)

type Error struct {
	errCode uint8
}

func (e Error) Error() string {
	panic("implement me")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestErrorNil(t *testing.T) {
	var e *Error
	var e1 Error
	fmt.Println(e, e1)
	//checkError(e)
}
