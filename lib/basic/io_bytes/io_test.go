package io_bytes

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestBufferIo(t *testing.T) {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

}

func TestReadAll(t *testing.T) {
	//r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	//b, err := io.ReadAll(r)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s", b)

	//r := strings.NewReader("some io.Reader stream to be read\n")
	//buf := make([]byte, 14)
	//if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", buf)
	//
	//// buffer smaller than minimal read size.
	//shortBuf := make([]byte, 3)
	//if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
	//	fmt.Println("error:", err)
	//}
	//
	//// minimal read size bigger than io.Reader stream
	//longBuf := make([]byte, 64)
	//if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
	//	fmt.Println("error:", err)
	//}

	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}
}

func TestPipe(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func TestCopy(t *testing.T) {
	//r := strings.NewReader("some io.Reader stream to be read\n")
	//size, err := io.Copy(os.Stdout, r)
	//if err != nil {
	//	fmt.Println("copy error cause: ", err)
	//	return
	//}
	//fmt.Println("copy size is ", size)

	//r1 := strings.NewReader("first reader\n")
	//r2 := strings.NewReader("second reader\n")
	//buf := make([]byte, 8)
	//
	//// buf is used here...
	//if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
	//	log.Fatal(err)
	//}
	//
	//// ... reused here also. No need to allocate an extra buffer.
	//if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
	//	log.Fatal(err)
	//}

	r := strings.NewReader("some io.Reader stream to be read")
	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}
}
