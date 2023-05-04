package manager

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserDao(t *testing.T) {
	convey.Convey("TestUserDao", t, func() {
		Convey("When ApplyFunc For Func", func() {
			// 支持为函数打一个桩
			want := "mock"
			selectOnePatch := gomonkey.ApplyFunc(SelectOne, func(id string) string {
				return want
			})
			defer selectOnePatch.Reset()
			ret := SelectOne("1")
			convey.So(ret, convey.ShouldEqual, want)
		})
	})
}
