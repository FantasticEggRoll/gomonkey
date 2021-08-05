package test

import (
	"errors"
	. "github.com/agiledragon/gomonkey/v2"
	"github.com/agiledragon/gomonkey/v2/test/abc"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var globalVar = "hello"

func TestGoMonkey(t *testing.T) {
	Convey("test apply", t, func() {
		Convey("test mock function", func() {
			patches := ApplyFunc(abc.HelloWorld, func() (error, string) {
				return nil, "mock function"
			})
			defer patches.Reset()
			_, result := abc.HelloWorld()
			So(result, ShouldEqual, "mock function")
		})

		Convey("test mock method", func() {
			test := abc.Test{A: 1, B: 2}
			fn1 := (*abc.Test).Hello
			fn2 := abc.Test.Hi
			patches := Apply(fn1, func(_ *abc.Test) int {
				return 0
			}).Apply(fn2, func(_ abc.Test) int {
				return 0
			})
			defer patches.Reset()
			result1 := test.Hello()
			result2 := test.Hi()
			So(result1, ShouldEqual, 0)
			So(result2, ShouldEqual, 0)
		})

		Convey("test mock global var", func() {
			patches := Apply(&globalVar, "mock global var")
			defer patches.Reset()
			So(globalVar, ShouldEqual, "mock global var")
		})

		Convey("test mock func var", func() {
			patches := Apply(abc.FnVal, func() int {
				return 3
			})
			defer patches.Reset()
			result := abc.FnVal()
			So(result, ShouldEqual, 3)
		})

		Convey("test mock func seq", func() {
			mockErr := errors.New("error")
			patches := Apply(abc.HelloWorld, []OutputCell{
				{
					Values: Params{nil, "a"},
				},
				{
					Values: Params{nil, "b"},
				},
				{
					Values: Params{mockErr, ""},
				},
			})
			defer patches.Reset()
			_, result := abc.HelloWorld()
			So(result, ShouldEqual, "a")
			_, result = abc.HelloWorld()
			So(result, ShouldEqual, "b")
			err, _ := abc.HelloWorld()
			So(err, ShouldEqual, mockErr)
		})

		Convey("test method seq", func() {
			test := abc.Test{A: 1, B: 2}
			fn1 := (*abc.Test).Hello
			fn2 := abc.Test.Hi
			patches := Apply(fn1, []OutputCell{
				{
					Values: Params{10},
				},
				{
					Values: Params{20},
				},
			}).Apply(fn2, []OutputCell{
				{
					Values: Params{30},
				},
				{
					Values: Params{40},
				},
			})
			defer patches.Reset()
			result := test.Hello()
			So(result, ShouldEqual, 10)
			result = test.Hello()
			So(result, ShouldEqual, 20)
			result = test.Hi()
			So(result, ShouldEqual, 30)
			result = test.Hi()
			So(result, ShouldEqual, 40)
		})

		Convey("test func var seq", func() {
			patches := Apply(abc.FnVal, []OutputCell{
				{
					Values: Params{50},
				},
				{
					Values: Params{60},
				},
			})
			defer patches.Reset()
			result := abc.FnVal()
			So(result, ShouldEqual, 50)
			result = abc.FnVal()
			So(result, ShouldEqual, 60)
		})

	})

}
