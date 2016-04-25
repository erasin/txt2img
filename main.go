package main

import (
	"html/template"

	"github.com/lunny/tango"
	//"github.com/unrolled/render"
)

func main() {

	tg := tango.Classic()

	tg.Get("/:to/:from/:date", new(TxtImgCon))

	//graceful.Run(":9092", 10*time.Second, st)
	tg.Run(":9092")
}

type TxtImgCon struct {
	tango.Params
	//tango.GZip
	tango.Ctx
}

//func (t TxtImgCon) Get(w http.ResponseWriter, req *http.Request) {
func (t TxtImgCon) Get() {

	StrTo := t.Params.Get(":to")
	StrFrom := t.Params.Get(":from")
	StrDate := t.Params.Get(":date")
	//w.Write([]byte(Name))
	//t.Ctx.ResponseWriter.Write([]byte(Name))
	//return

	var tplMain = `尊敬的 {{.To}}:

	这里是您的邀请函,你可以凭借此文件来我公司进行商谈.

									重要: {{.From}}
									{{.Date}}

`
	tc := &TextConvert{font: "/Library/Fonts/华文仿宋.ttf"}

	tpMain, _ := template.New("main").Parse(tplMain)
	tpMain.Execute(tc, map[string]string{"To": StrTo, "From": StrFrom, "Date": StrDate})

	//bf := bufio.NewWriter(outFile)
	//var b []byte
	//bf := bufio.NewBufferString(&b)
	//tpMain.Execute(bf, map[string]string{"To": StrTo, "From": StrFrom, "Date": StrDate})
	//bf.Read()

	tc.doImg()
	tc.writeTo(t.Ctx.ResponseWriter)
	//return

}
