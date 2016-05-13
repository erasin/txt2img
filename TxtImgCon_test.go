package main

import (
	"reflect"
	"testing"

	"github.com/lunny/tango"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTxtImgConGet(t *testing.T) {
	//req := TxtImgCon{}

	tango.Context{}

	Convey("获取图像:", t, func() {

		a, b := "str", "str2"
		_ := reflect.DeepEqual(a, b)

	})
}
