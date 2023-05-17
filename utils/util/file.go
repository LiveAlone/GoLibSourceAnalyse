package util

import (
	"bufio"
	"errors"
	"os"
)

func ReadFileLines(filePath string) (fileLines []string, err error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func WriteFileLines(path string, lines []string) error {
	data := make([]byte, 0)
	for _, line := range lines {
		data = append(data, line...)
		data = append(data, '\n')
	}
	return os.WriteFile(path, data, 0644)
}

func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func CreateDirIfNotExists(dirPath string) error {
	fileInfo, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// 创建文件
		return os.Mkdir(dirPath, 0755)
	}

	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		return nil
	}
	return errors.New("file exists but not dir")
}
