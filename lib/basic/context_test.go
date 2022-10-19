package basic

import (
	"context"
	"fmt"
	"testing"
)

type LocalTestDoneCtx struct {
	context.Context
	Name  string
}

func TestLocalContext(t *testing.T) {
	l := LocalTestDoneCtx{Context: context.TODO(), Name: "123"}
	fmt.Println(l.Value(123))
}
