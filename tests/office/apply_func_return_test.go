package test

import (
	"github.com/agiledragon/gomonkey/v2/test/fake"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

/*
compare with apply_func_seq_test.go
*/
func TestApplyFuncReturn(t *testing.T) {
	Convey("TestApplyFuncReturn", t, func() {

		Convey("declares the values to be returned", func() {
			info1 := "hello cpp"

			patches := ApplyFuncReturn(fake.ReadLeaf, info1, nil)
			defer patches.Reset()

			for i := 0; i < 10; i++ {
				output, err := fake.ReadLeaf("")
				So(err, ShouldEqual, nil)
				So(output, ShouldEqual, info1)
			}

			patches.Reset() // if not reset will occur:patch has been existed
			info2 := "hello golang"
			patches.ApplyFuncReturn(fake.ReadLeaf, info2, nil)
			for i := 0; i < 10; i++ {
				output, err := fake.ReadLeaf("")
				So(err, ShouldEqual, nil)
				So(output, ShouldEqual, info2)
			}
		})
	})
}
