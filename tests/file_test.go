package tests

import (
	"fmt"
	"testing"

	"avenger/tools/file"
)

func TestFileReadLines(t *testing.T) {
	path := "sample.txt"
	result, err := file.ReadLines(path, 5, 2)
	testingReport(t, result, err)
	result, err = file.ReadLines(path, 0, 1)
	testingReport(t, result, err)
	result, err = file.ReadLines(path, 5, 0)
	testingReport(t, result, err)
}

func TestFileMkdir(t *testing.T) {
	fmt.Println(file.Mkdir("test"))
}

func TestFileCheck(t *testing.T) {
	//文件夹
	fmt.Println("test is empty:", file.DirExists("test"))
	fmt.Println("test1 is empty:", file.DirExists("test1"))
	fmt.Println("../tests is empty:", file.DirExists("../tests"))
	//文件都生效
	fmt.Println("sample.txt is exists:", file.Exists("./sample/sample.txt"))
	fmt.Println("sample1.txt is exists:", file.Exists("./sample/sample2.txt"))
	fmt.Println("test is exists:", file.Exists("test"))
	fmt.Println("test1 is empty:", file.DirExists("test1"))
	//文件夹
	fmt.Println("test is empty:", file.IsEmpty("test"))
	fmt.Println("test1 is empty:", file.IsEmpty("test1"))
	fmt.Println("../tests is empty:", file.IsEmpty("../tests"))
	//文件
	fmt.Println("test.txt is empty:", file.IsEmpty("test.txt"))
	fmt.Println("test1.txt is empty:", file.IsEmpty("test1.txt"))
	fmt.Println("sample.txt is empty:", file.IsEmpty("./sample/sample.txt"))
}

func TestFileRemove(t *testing.T) {
	fmt.Println(file.RemoveAll("test"))
}

func TestFileCopy(t *testing.T) {
	fmt.Println(file.CopyTo("./sample/sample.txt", "./sample/sample1.txt"))
}
