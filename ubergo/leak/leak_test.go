package leak

import (
	"go.uber.org/goleak"
	"testing"
)

func TestFirst(t *testing.T) {
	defer goleak.VerifyNone(t)

	// test logic here.
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
