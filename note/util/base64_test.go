package util

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestBase64(t *testing.T) {
	//读文件
	ff, _ := os.Open("C:\\Users\\71013\\Desktop\\pdf-test.pdf")
	fileInfo, _ := ff.Stat()
	defer ff.Close()
	sourceBuffer := make([]byte, fileInfo.Size())
	n, _ := ff.Read(sourceBuffer)
	//base64s
	sourceString := base64.StdEncoding.EncodeToString(sourceBuffer[:n])
	t.Logf(sourceString)
}
