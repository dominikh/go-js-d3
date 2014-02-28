package d3

import (
	"net/url"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/util"
)

var d3 = js.Global.Get("d3")

type Selection struct {
	js.Object
}

func Select(sel string) Selection {
	return Selection{d3.Call("select", sel)}
}

func SelectAll(sel string) Selection {
	return Selection{d3.Call("selectAll", sel)}
}

func (s Selection) Append(name string) Selection {
	return Selection{s.Call("append", name)}
}

func (s Selection) Attr(name, value string) Selection {
	return Selection{s.Call("attr", name, value)}
}

func (s Selection) Select(sel string) Selection {
	return Selection{s.Call("select", sel)}
}

func (s Selection) SelectAll(sel string) Selection {
	return Selection{s.Call("selectAll", sel)}
}

func (s Selection) Enter() Selection {
	return Selection{s.Call("enter")}
}

func (s Selection) SetStyle(name, value string) {
	if len(value) == 0 {
		s.Call("style", name, nil)
	} else {
		s.Call("style", name, value)
	}
}

type Data interface {
	Underlying() js.Object
}

func (sel Selection) SetData(data []Data) Selection {
	mapped := make([]js.Object, len(data))
	for i, e := range data {
		mapped[i] = e.Underlying()
	}

	return Selection{sel.Call("data", mapped)}
}

func JSON(url *url.URL, fn func(err error, data js.Object)) {
	d3.Call("json", url.String(), func(err js.Object, data js.Object) {
		fn(util.Error(err), data)
	})
}

type Formatter func(js.Object) string

func Format(spec string) Formatter {
	f := d3.Call("format", spec)
	return func(o js.Object) string { return f.Invoke(o).String() }
}
