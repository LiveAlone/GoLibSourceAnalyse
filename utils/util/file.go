package util

import (
	"bufio"
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

func WriteFileLines(path string, lines []string) error {
	data := make([]byte, 0)
	for _, line := range lines {
		data = append(data, line...)
	}
	return os.WriteFile(path, data, 0644)
}
