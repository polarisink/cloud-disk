package test

import (
	"cloud-disk/core/helper"
	"os"
	"testing"
)

func TestOssUpload(t *testing.T) {
	file,err:=os.Open("/Users/lqs/Desktop/resume.jpg")
	if err!=nil {
		t.Fatal("该文件不存在")
	}
	upload, err := helper.OssLocalFile(file)
	if err != nil {
		return
	}
	t.Log(upload)
}
