package dao

import (
	"encoding/json"
	"testing"
)

func TestInsertBook(t *testing.T) {
	//for i := 0; i < 1; i++ {
	//	book := &model.Book{Name: "test"}
	//	InsertBook(book)
	//}
}

func TestDeleteBookByName(t *testing.T) {
	//num , err :=DeleteBookByName("test")
	//if num==0 || err!=nil{
	//	t.Error("")
	//}
}

func TestListBooksByCreatedTime(t *testing.T) {
	books := ListBooksByCreatedTime(0, 0, -1)
	bytes, _ := json.Marshal(books)

	t.Logf("books = %v", string(bytes))
}
