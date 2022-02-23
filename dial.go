package main

import (
	"fmt"

	"github.com/r0123r/go-iup/iup"
)

func vbox5(name string) *iup.Handle {
	txtX := iup.Text("READONLY=YES,MINSIZE=150x0")
	txtY := iup.Text("READONLY=YES,MINSIZE=150x0")
	res := iup.Vbox(
		"TABTITLE="+name,
		iup.Fill(),
		iup.Label("TITLE=\"Угол по X\""), txtX,
		iup.Dial(func(arg *iup.DialMouseMove) {
			txtX.SetAttribute("VALUE", fmt.Sprint(arg.Angle))
			arg.Return = iup.DEFAULT
		}),
		iup.Label("TITLE=\"Угол по Y\""), txtY,
		iup.Dial("ORIENTATION=VERTICAL", func(arg *iup.DialMouseMove) {
			txtY.SetAttribute("VALUE", fmt.Sprint(arg.Angle))
			arg.Return = iup.DEFAULT
		}),
		iup.Fill(),
	)
	return res
}
