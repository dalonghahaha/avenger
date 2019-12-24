package file

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

const (
	READOYLY  = 0644
	READWRITE = 0755
)

var fileSystem = afero.NewOsFs()

func AbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}

func IsDir(path string) bool {
	result, err := afero.IsDir(fileSystem, path)
	if err != nil {
		return false
	}
	return result
}

func IsEmpty(path string) bool {
	result, err := afero.IsEmpty(fileSystem, path)
	if err != nil {
		return true
	}
	return result
}

func DirExists(path string) bool {
	result, err := afero.DirExists(fileSystem, path)
	if err != nil {
		return false
	}
	return result
}

func Exists(path string) bool {
	result, err := afero.Exists(fileSystem, path)
	if err != nil {
		return false
	}
	return result
}

func Mkdir(path string) error {
	return fileSystem.MkdirAll(path, READWRITE)
}

func Remove(path string) error {
	return fileSystem.Remove(path)
}

func RemoveAll(path string) error {
	return fileSystem.RemoveAll(path)
}

func CopyTo(src string, dst string) error {
	srcFile, err := fileSystem.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		srcFile.Close()
	}()
	destFile, err := fileSystem.OpenFile(dst, os.O_WRONLY|os.O_CREATE, READWRITE)
	if err != nil {
		return err
	}
	defer func() {
		destFile.Close()
	}()
	_, err = io.Copy(destFile, srcFile)
	return err
}

func Read(path string) ([]byte, error) {
	return afero.ReadFile(fileSystem, path)
}

func ReadDir(path string) ([]os.FileInfo, error) {
	return afero.ReadDir(fileSystem, path)
}

func Write(path string, content []byte) error {
	return afero.WriteFile(fileSystem, path, content, READWRITE)
}

func Append(path string, content []byte) error {
	_file, err := fileSystem.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, READWRITE)
	defer func() {
		_file.Close()
	}()
	if err != nil {
		return err
	}
	_, err = _file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func ReadText(path string) (string, error) {
	data, err := Read(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteText(path string, content string) error {
	return Write(path, []byte(content))
}

func AppendText(path string, content string) error {
	return Append(path, []byte(content))
}

func ReadLines(path string, offset int, limit int) (string, error) {
	f, err := fileSystem.Open(path)
	if err != nil {
		return "", err
	}
	result := ""
	buf := bufio.NewReader(f)
	for i := offset; i < offset+limit; i++ {
		line, err := buf.ReadString('\n')
		if i >= offset {
			result = result + line
		}
		if err != nil {
			if err == io.EOF {
				return result, nil
			}
			return "", err
		}
	}
	return result, nil
}
