package type_example

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap(t *testing.T) {
	err := fmt.Errorf("test %d", 123)
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
}
