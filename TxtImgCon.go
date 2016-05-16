package main

import (
	"fmt"
	"net/http"
)

//func (t TxtImgCon) Get(w http.ResponseWriter, req *http.Request) {

func TxtImgConGet(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	// 测试主体
	var tplMain = `尊敬的 {{.To}}:

	这里是您的邀请函,你可以凭借此文件来我公司进行商谈.

									重要: {{.From}}
									{{.Date}}

`
	var fields = map[string]string{"To": "", "Form": "", "Date": "222"}

	//fontPath := "/Library/Fonts/华文仿宋.ttf"
	fontPath := "华文仿宋.ttf"

	tc, err := NewTextConvert(fontPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = NewTpl(tplMain, fields).Encoder(req.Form, tc)
	if err != nil {
		fmt.Println(err.Error())
	}
	tc.EncodeImg().writeTo(w)
}
