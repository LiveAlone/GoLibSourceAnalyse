package basic

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestLogBasic(t *testing.T) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
	fmt.Println(&buf)
}
