package gobase

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FileOrFolderExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			// return some error
			return false
		}
	}
	return true
}

func IsFolder(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// err标识：1：读取失败，2：文件不存在
func ReadFileAsString(filePath string) (*string, int) {
	if !FileOrFolderExist(filePath) {
		// file not exist
		return nil, 2
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, 1
	}

	result := string(b)

	return &result, 0

}

func ReadFileToLines(filePath string) ([]string, int) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		log.Println("error", err)
		return nil, 1
	}

	defer file.Close()
	result := make([]string, 0)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("read string error", err)
			break
		}

		line = strings.TrimSpace(line)
		//if len(line) > 0 {
		result = append(result, line)
		//}
	}

	return result, 0
}

func WriteStringToFile(content, fileName string, overWriteIfExist bool) bool {
	if FileOrFolderExist(fileName) {
		if overWriteIfExist {
			os.Remove(fileName)
		} else {
			return false
		}
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		log.Println("error", err)
		return false
	}

	defer file.Close()

	fileWrite := bufio.NewWriter(file)
	fileWrite.WriteString(content)

	fileWrite.Flush()

	return true
}
