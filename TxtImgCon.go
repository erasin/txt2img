package main

import "github.com/lunny/tango"

type TxtImgCon struct {
	tango.Params
	//tango.GZip
	tango.Ctx
	tango.Logger
}

//func (t TxtImgCon) Get(w http.ResponseWriter, req *http.Request) {
func (t TxtImgCon) Get() {

	// 测试主体
	var tplMain = `尊敬的 {{.To}}:

	这里是您的邀请函,你可以凭借此文件来我公司进行商谈.

									重要: {{.From}}
									{{.Date}}

`
	var fields = map[string]string{"To": "", "Form": "", "Date": "222"}

	fontPath := "/Library/Fonts/华文仿宋.ttf"
	tc, err := NewTextConvert(fontPath)
	if err != nil {
		t.Error(err.Error())
	}
	err = NewTpl(tplMain, fields).Encoder(t.Forms(), tc)
	if err != nil {
		t.Error(err.Error())
	}
	tc.EncodeImg().writeTo(t.Ctx.ResponseWriter)
}
