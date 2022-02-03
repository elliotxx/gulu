package os

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}

	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !s.IsDir()
}

// 按行读取文本文件
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// 字符串截取
func substr(s string, pos, length int) string {
	runes := []rune(s)

	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}

	return string(runes[pos:l])
}

// 获取上一级目录
func GetParentDir(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// 获取 go mod 项目的根目录，
// 从当前目录依次上下查找，知道找到当前目录中存在 go.mod 文件的目录，即根目录
func GetGoModRootDir(curDir string) (string, error) {
	files, err := ioutil.ReadDir(curDir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() && file.Name() == "go.mod" {
			return curDir, nil
		}
	}

	rootDir, err := GetGoModRootDir(GetParentDir(curDir))
	if err != nil {
		return "", err
	}

	return rootDir, nil
}
