package gopatch

import (
	"errors"
	"fmt"
	"testing"
)

func TestNone(t *testing.T) {
	err := errors.New(fmt.Sprintf("hello"))
	fmt.Println(err)
}
