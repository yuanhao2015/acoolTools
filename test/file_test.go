package test

import (
	"bufio"
	"fmt"
	"github.com/yuanhao2015/acoolTools"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	//判断文件或目录是否存在
	exists := acoolTools.FileUtils.Exists("F:/go-test/test")
	fmt.Println("创建前------------------------>", exists)
	err := os.MkdirAll("F:/go-test/test", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	exists = acoolTools.FileUtils.Exists("F:/go-test/test")
	fmt.Println("创建后------------------------>", exists)
}

func TestIsDir(t *testing.T) {
	//判断是否是目录
	dir := acoolTools.FileUtils.IsDir("F:/go-test/test")
	fmt.Println("是否是目录--------------------->", dir)
	dir = acoolTools.FileUtils.IsDir("F:/go-test/test/test.txt")
	fmt.Println("是否是目录--------------------->", dir)
}

func TestIsFile(t *testing.T) {
	//判断是否是文件
	file := acoolTools.FileUtils.IsFile("F:/go-test/test")
	fmt.Println("是否是文件--------------------->", file)
	file = acoolTools.FileUtils.IsFile("F:/go-test/test/test.txt")
	fmt.Println("是否是文件--------------------->", file)
}

func TestRemove(t *testing.T) {
	//删除文件
	file, err := acoolTools.FileUtils.RemoveFile("F:/go-test/test/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if file {
		//查看文件是否还存在
		exists := acoolTools.FileUtils.Exists("F:/go-test/test/test.txt")
		fmt.Println("文件是否存在------------------------>", exists)
	}
}

func TestOpenFileWronly(t *testing.T) {
	//用只写模式打开一个文件，并且写入5条内容,若文件不存在将会创建一个
	path := "F:/go-test/test/test.txt"
	str := "hello word acoolTools \n"
	wronly, err := acoolTools.FileUtils.OpenFileWronly(path)
	if err != nil {
		t.Fatal(err)
	}
	defer wronly.Close()
	write := bufio.NewWriter(wronly)
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	//读取文件写入到控制台
	files, err := acoolTools.FileUtils.OpenFileRdonly(path)
	if err != nil {
		t.Fatal(err)
	}
	defer files.Close()
	reader := bufio.NewReader(files)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(str)
	}
}

func TestOpenFileRdonly(t *testing.T) {
	path := "F:/go-test/test/test.txt"
	files, err := acoolTools.FileUtils.OpenFileRdonly(path)
	if err != nil {
		t.Fatal(err)
	}
	defer files.Close()
	reader := bufio.NewReader(files)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(str)
	}
}

func TestOpenFileAppend(t *testing.T) {
	//打开文件在文件后面追加数据
	path := "F:/go-test/test/test.txt"
	str := "追加内容 \n"
	wronly, err := acoolTools.FileUtils.OpenFileAppend(path)
	if err != nil {
		t.Fatal(err)
	}
	defer wronly.Close()
	write := bufio.NewWriter(wronly)
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
	//读取文件写入到控制台
	files, err := acoolTools.FileUtils.OpenFileRdonly(path)
	if err != nil {
		t.Fatal(err)
	}
	defer files.Close()
	reader := bufio.NewReader(files)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(str)
	}
}

func TestFileCopy(t *testing.T) {
	//文件复制功能
	path := "F:/go-test/test/test.txt"
	copyPath := "F:/go-test/test/test.txt1"
	//复制钱
	exists := acoolTools.FileUtils.Exists(copyPath)
	fmt.Println("复制前文件是否存在------------------>", exists)
	//复制后
	fileCopy, err := acoolTools.FileUtils.FileCopy(path, copyPath)
	if err != nil {
		t.Fatal(err)
	}
	if fileCopy {
		exists := acoolTools.FileUtils.Exists(copyPath)
		fmt.Println("复制前文件是否存在------------------>", exists)
	}
}
