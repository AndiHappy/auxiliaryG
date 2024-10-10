package util

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

// CreateDirectory 创建文件夹
func CreateDirectory(fP string) bool {
	_, err := os.Stat("test")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	return true
}

// CreateFile 创建文件：如果文件存在，则返回对应的文件，如果文件不存在则创建
func CreateFile(fN string) (*os.File, error) {
	return os.OpenFile(fN, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(fP string) bool {
	_, err := os.Stat(fP) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// DeleteFile 删除文件
func DeleteFile(fp string) bool {
	err := os.Remove(fp)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// ReadFileByLineV1 一行一行的读取文件内容
func ReadFileByLineV1(fP string) ([]string, error) {
	res := make([]string, 0)
	fileBuffer, err := os.ReadFile(fP)
	if err != nil {
		return nil, err
	}
	data := bufio.NewScanner(bytes.NewReader(fileBuffer))
	data.Split(bufio.ScanRunes)
	for data.Scan() {
		res = append(res, data.Text())
	}
	return res, nil
}

// ReadFileByLineV2 ...
func ReadFileByLineV2(fName string) ([]string, error) {
	file, err := os.Open(fName)
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

// AppendFileContent ... 项文件中增加内容
func AppendFileContent(content, fName string) (int, error) {
	f, err := os.OpenFile(fName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		CheckErr(err)
		return 0, err
	}
	defer f.Close()
	n, err := f.WriteString(content)
	if err != nil {
		CheckErr(err)
		return 0, err
	} else {
		return n, nil
	}
	return 0, nil
}
