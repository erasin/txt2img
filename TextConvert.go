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

type TextConvert struct {
	Text string
	rgba *image.RGBA
	png  string
	font string
}

type textLine struct {
	Text  string
	Index int
}

// 字符串拆解
func (tc *TextConvert) sliptString(lines *list.List, lineSize int) (countAll int) {
	log.Println(tc.Text)
	countAll = 0
	text := strings.Replace(tc.Text, "\t", "    ", 0)
	texts := strings.Split(text, "\n")

	for _, v := range texts {

		//c := runewidth.StringWidth(v)

		txt := []rune(v)
		// 处理回车和换行
		c := len(txt)
		var count int

		//log.Println(c/lineSize)

		if n := c / lineSize; n <= 0 {
			count = 1
		} else {
			count = n + 1
		}

		//log.Println(c, count)

		for j := 0; j < count; j++ {
			start := j * lineSize
			end := (j + 1) * lineSize
			if j == count-1 {
				//log.Println(start, end,countAll+j)
				lines.PushBack(textLine{string(txt[start:]), countAll + j})
				// t.Lines[j] = string(txt[start:])
			} else {
				//log.Println(start, end,countAll+j)
				lines.PushBack(textLine{string(txt[start:end]), countAll + j})
				// t.Lines[j] = string(txt[start:end])
			}
		}

		countAll += count
		//countAll += 1
		//lines.PushBack(textLine{"",countAll})
	}

	return
}

func (tc *TextConvert) wrap(lines *list.List, lineSize int) (countAll int) {

	text := WrapString(tc.Text, uint(lineSize))
	texts := strings.Split(text, "\n")
	countAll = len(texts)
	for k, v := range texts {
		lines.PushBack(textLine{v, k})
	}
	return
}

func (tc *TextConvert) Write(p []byte) (n int, err error) {
	tc.Text += string(p)
	return 0, nil
}

func (tc *TextConvert) doImg() {
	var size = 25.0
	var dx = 500
	var dy = 300

	lineSize := int(dx / int(size))

	// 计算行数以及拆解字符串
	lines := list.New()
	//tl := tc.sliptString(lines, lineSize)
	tl := tc.wrap(lines, lineSize)
	dy = tl * int(size)

	fontb, err := ioutil.ReadFile(tc.font)
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
		pt := freetype.Pt(0, (t.Index+1)*intSize)
		c.DrawString(t.Text, pt)
	}

	return
}

func (tc *TextConvert) writeTo(w io.Writer) {
	//bf := bufio.NewWriter(w)
	//err := png.Encode(w, tc.rgba)
	err := jpeg.Encode(w, tc.rgba, &jpeg.Options{80})

	//err = png.Encode(bf, tc.rgba)
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
