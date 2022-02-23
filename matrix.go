package main

import (
	"github.com/r0123r/go-iup/iup"
)

func vbox7(name string) *iup.Handle {
	mtx := iup.Matrix(
		"NUMCOL=5,NUMLIN=10,NUMCOL_VISIBLE=5,NUMLIN_VISIBLE=3,WIDTHDEF=34",
		"RESIZEMATRIX=YES",
	)
	mtx.SetAttrs("1:1", "test", "1:2", "A", "1:3", "B", "1:4", "C", "1:5", "D")
	mtx.SetAttribute("2:2", "5.6\n3.33")
	mtx.SetAttribute("HEIGHT2", "30")
	mtx.SetAttribute("WIDTH2", "190")

	mtx.SetAttribute("BGCOLOR1:*", "192 192 192")
	mtx.SetAttribute("SORTSIGN1", "DOWN")
	mtx.SetAttribute("TYPE4:1", "COLOR")
	mtx.SetAttribute("4:1", "255 0 128")
	mtx.SetAttribute("TYPE4:2", "FILL")
	mtx.SetAttribute("4:2", "60")
	mtx.SetAttribute("SHOWFILLVALUE", "Yes")

	mtx.SetCallback(func(arg *iup.MatrixDropCheck) {
		if arg.Lin == 4 && arg.Col == 4 {
			arg.Return = iup.CONTINUE
		} else {
			arg.Return = iup.IGNORE
		}
	})
	mtx.SetCallback(func(arg *iup.MatrixDrop) {
		if arg.Lin == 3 && arg.Col == 1 {
			arg.Drop.SetAttribute("1", "A - Test of Very Big String for Dropdown!")
			// arg.Drop.SetAttribute("2", "B")
			// arg.Drop.SetAttribute("3", "C")
			// arg.Drop.SetAttribute("4", "XXX")
			// arg.Drop.SetAttribute("5", "5")
			// arg.Drop.SetAttribute("6", "6")
			// arg.Drop.SetAttribute("7", "7")
			// arg.Drop.SetAttribute("8", "")
			// arg.Drop.SetAttribute("VALUE", "4")
			arg.Return = iup.DEFAULT
		} else {

			arg.Return = iup.IGNORE
		}
	})
	mtx.SetAttribute("TOGGLEVALUE4:4", "ON")
	mtx.SetAttribute("4:4", "1")
	mtx.SetAttribute("TOGGLECENTERED", "Yes")

	res := iup.Vbox(
		"TABTITLE="+name,
		mtx,
	)
	return res
}
