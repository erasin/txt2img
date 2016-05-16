package main

import (
	"bufio"
	"container/list"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	//"github.com/unrolled/render"
)

// TextConvert 文本转化
type TextConvert struct {
	Text string
	rgba *image.RGBA
	png  string
	Font string
}

// textLine 单行数据
type textLine struct {
	Text  string
	Index int
}

// NewTextConvert 新建转换主体
func NewTextConvert(ft string) (*TextConvert, error) {
	_, err := os.Stat(ft)
	if err != nil {
		// 获取默认字体,打印输出错误
		return &TextConvert{}, err
	}
	return &TextConvert{Font: ft}, nil
}

// sliptString 字符串拆解
func (tc *TextConvert) sliptString(lines *list.List, lineSize int) (countAll int) {
	//log.Println(tc.Text)
	countAll = 0
	text := strings.Replace(tc.Text, "\t", "    ", 0)
	texts := strings.Split(text, "\n")

	for _, v := range texts {

		txt := []rune(v)
		// 处理回车和换行
		c := len(txt)
		var count int

		if n := c / lineSize; n <= 0 {
			count = 1
		} else {
			count = n + 1
		}

		for j := 0; j < count; j++ {
			start := j * lineSize
			end := (j + 1) * lineSize
			if j == count-1 {
				lines.PushBack(textLine{string(txt[start:]), countAll + j})
			} else {
				lines.PushBack(textLine{string(txt[start:end]), countAll + j})
			}
		}

		countAll += count
	}

	return
}

// wrap 字符换行处理
func (tc *TextConvert) wrap(lines *list.List, lineSize int) (countAll int) {
	text := WrapString(tc.Text, uint(lineSize))
	texts := strings.Split(text, "\n")
	countAll = len(texts)
	for k, v := range texts {
		lines.PushBack(textLine{v, k})
	}
	return
}

// Write 实现 io.Writer 接口 写入文字实体
func (tc *TextConvert) Write(p []byte) (n int, err error) {
	tc.Text += string(p)
	return 0, nil
}

// EncodeImg 处理图片生成
func (tc *TextConvert) EncodeImg() *TextConvert {
	var size = 25.0
	var dx = 1500
	var dy = 1300

	lineSize := int(dx / int(size))

	// 计算行数以及拆解字符串
	lines := list.New()
	//tl := tc.sliptString(lines, lineSize)
	tl := tc.wrap(lines, lineSize)
	dy = tl * int(size)

	fontb, err := ioutil.ReadFile(tc.Font)
	if err != nil {
		log.Println(err)
	}
	fontf, err := truetype.Parse(fontb)
	if err != nil {
		log.Println(err)
	}
	fg, bg := image.Black, image.White
	tc.rgba = image.NewRGBA(image.Rect(0, 0, dx, dy))
	draw.Draw(tc.rgba, tc.rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(fontf)
	c.SetFontSize(size)
	c.SetClip(tc.rgba.Bounds())
	c.SetDst(tc.rgba)
	c.SetSrc(fg)

	opts := truetype.Options{}
	opts.Size = size
	// face := truetype.NewFace(f, &opts)
	// faceWidth,ok := face.GlyphAdvance(rune("中"))
	intSize := int(size)

	for txt := lines.Front(); txt != nil; txt = txt.Next() {
		t := txt.Value.(textLine)
		pt := freetype.Pt(0, (t.Index+2)*intSize)
		c.DrawString(t.Text, pt)
	}

	return tc
}

func (tc *TextConvert) writeTo(w io.Writer) {
	//bf := bufio.NewWriter(w)
	//err := png.Encode(w, tc.rgba)
	err := jpeg.Encode(w, tc.rgba, &jpeg.Options{80})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (tc *TextConvert) saveImg() {
	outFile, _ := os.Create("demo2.png")
	defer outFile.Close()
	bf := bufio.NewWriter(outFile)
	err := png.Encode(bf, tc.rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = bf.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
