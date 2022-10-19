package eddycjy

import (
	"fmt"
	"testing"
)

type MyError struct {
	Msg string
}

func (m *MyError) Error() string {
	return m.Msg
}

func TestError(t *testing.T) {
	var err error
	c := GetErr()
	fmt.Println(c == nil)
	err = GetErr()
	fmt.Println(err == nil)
}

func GetErr() *MyError{
	return nil
}