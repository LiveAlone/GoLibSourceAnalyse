package model

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestUserDao(t *testing.T) {
	Convey("TestUserDao", t, func() {
		Convey("When ApplyFunc For Func", func() {
			// 支持为函数打一个桩
			want := "mock"
			selectOnePatch := gomonkey.ApplyFunc(SelectOne, func(id string) string {
				return want
			})
			defer selectOnePatch.Reset()
			ret := SelectOne("1")
			So(ret, ShouldEqual, want)
		})

		// 支持为函数打一个特定的桩序列
		Convey("When ApplyFuncSeq For Func", func() {
			want := []gomonkey.OutputCell{
				{Values: gomonkey.Params{"1"}},
				{Values: gomonkey.Params{"2"}},
				{Values: gomonkey.Params{"3"}},
			}
			selectList := gomonkey.ApplyFuncSeq(SelectList, want)
			defer selectList.Reset()
			Convey("Then Test ApplyFuncSeq Patch", func() {
				So(SelectList(), ShouldEqual, "1")
				So(SelectList(), ShouldEqual, "2")
				So(SelectList(), ShouldEqual, "3")
			})
		})

		//支持为成员方法打一个桩
		Convey("When ApplyMethod For userDao Func", func() {
			u := &UserDao{}
			updateF := gomonkey.ApplyMethod(reflect.TypeOf(u), "Update", func(_ *UserDao, id, name, phoneNumber string) int64 {
				fmt.Println(id, name, phoneNumber)
				return 1
			})
			defer updateF.Reset()
			Convey("Then Test ApplyMethod userDao Func Patch", func() {
				So(u.Update("", "", ""), ShouldEqual, int64(1))
			})
		})

		//支持为全局变量打一个桩
		Convey("When ApplyGlobalVar", func() {
			globalVar := gomonkey.ApplyGlobalVar(&num, 10)
			defer globalVar.Reset()
			Convey("Then Test ApplyGlobalVar Patch", func() {
				So(num, ShouldEqual, int64(10))
			})
		})

		// 定义函数变量f1， f2
		f1 := func() string {
			return "f1"
		}
		f2 := func() string {
			return "f2"
		}
		// 支持为函数变量打桩
		Convey("When ApplyFuncVarSeq", func() {
			want := []gomonkey.OutputCell{
				{Values: gomonkey.Params{"1"}},
				{Values: gomonkey.Params{"2"}},
				{Values: gomonkey.Params{"3"}},
			}

			// 支持为函数变量打一个特定的桩序列
			funcVarSeq := gomonkey.ApplyFuncVarSeq(&f1, want)
			defer funcVarSeq.Reset()

			// 支持为函数变量打一个桩
			objVar := gomonkey.ApplyFuncVar(&f2, func() string {
				return "ggr"
			})
			defer objVar.Reset()

			Convey("Then Test ApplyFuncVarSeq&ApplyFuncVar Patch", func() {
				So(f1(), ShouldEqual, "1")
				So(f1(), ShouldEqual, "2")
				So(f1(), ShouldEqual, "3")
				So(f2(), ShouldEqual, "ggr")
			})
		})
	})
}
