package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

//分片大小
const chunkSize = 100 * 1024 * 1024
const dir = "/Users/lqs/Documents/CS学习/计算机操作系统/书籍/"
const lol = dir + "深入理解计算机系统.pdf"
const perm = 0666
const mode = os.O_CREATE | os.O_WRONLY
const appendMode = mode | os.O_APPEND

//文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat(lol)
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := int(fileInfo.Size() / chunkSize)
	myFile, err := os.OpenFile(lol, os.O_RDONLY, perm)
	if err != nil {
		t.Fatal(err)
	}
	defer myFile.Close()
	b := make([]byte, chunkSize)
	for i := 0; i < chunkNum+1; i++ {
		//指定读取文件的起始位置
		_, err := myFile.Seek(int64(i*chunkSize), 0)
		if err != nil {
			t.Fatal(err)
		}
		i2 := fileInfo.Size() - int64(i*chunkSize)
		if chunkSize > i2 {
			b = make([]byte, i2)
		}
		myFile.Read(b)
		err = ioutil.WriteFile(fmt.Sprintf("chunk_%d.mp4", i), b, perm)
		if err != nil {
			t.Fatal(err)
		}
		f, err := os.OpenFile(dir+strconv.Itoa(i)+".chunk", mode, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
}

// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile(lol, appendMode, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat(lol)
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile(dir+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 文件一致性校验
func TestCheck(t *testing.T) {
	// 获取第一个文件的信息
	file1, err := os.OpenFile(lol, os.O_RDONLY, perm)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	// 获取第二个文件的信息
	file2, err := os.OpenFile(lol, os.O_RDONLY, perm)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}

