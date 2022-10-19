package io_bytes

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	//var b bytes.Buffer // A Buffer needs no initialization.
	//b.Write([]byte("Hello "))
	//fmt.Fprintf(&b, "world!")
	//b.WriteTo(os.Stdout)
	//fmt.Println(b.Bytes())

	//var a, b []byte
	//if bytes.Compare(a, b) < 0 {
	//	// a less b
	//}

	//fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2))

	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
}
