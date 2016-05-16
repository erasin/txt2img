package main

import (
	"fmt"
	"html/template"
	"io"
	"net/url"
	"strings"
)

type Tpl struct {
	Tpl    string
	Fields map[string]string
}

// NewTpl 创建模板
func NewTpl(t string, fs map[string]string) *Tpl {
	return &Tpl{t, fs}
}

func (t *Tpl) Encoder(f url.Values, w io.Writer) error {
	if len(t.Fields) > 0 {
		for k, _ := range t.Fields {
			t.Fields[k] = f.Get(strings.ToLower(k))
		}
	}
	fmt.Println(t.Fields)

	tpMain, err := template.New("main").Parse(t.Tpl)
	if err != nil {
		return err
	}
	err = tpMain.Execute(w, t.Fields)
	if err != nil {
		return err
	}
	return nil
}
