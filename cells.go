package main

import (
	"fmt"
	"math"

	"github.com/r0123r/go-iup/iup"
)

var mlc iup.CellsMouseMotion
var scale int32
var fscale, fclick bool

func vbox71(name string) *iup.Handle {
	scale = 10
	mtx := iup.Matrix(
		"NUMCOL=6,NUMLIN=1,WIDTHDEF=34", "RASTERSIZE=x25", "READONLY=YES", "EXPAND=HORIZONTAL",
	)
	mtx.SetAttrs("1:1", "Col", "1:4", "Lin", "TYPE1:3", "FILL", "TYPE1:6", "FILL", "SHOWFILLVALUE", "Yes")
	cells := iup.Cells("BOXED=NO", "RASTERSIZE=40x15", iup.Attr("FONT", "Tahoma, 8"),
		func(arg *iup.CellsMouseMotion) {
			mlc.Column = arg.Column
			mlc.Line = arg.Line
		},
		func(arg *iup.CellsMouseClick) {
			mlc.Column = arg.Column
			mlc.Line = arg.Line
			fclick = true
		},
		func(arg *iup.CellsDraw) {
			cd := &iup.CD{arg.Cnv}
			var v [3]byte
			if math.Mod(float64(arg.Line), 2) == 0 && math.Mod(float64(arg.Column), 2) == 0 ||
				math.Mod(float64(arg.Line+1), 2) == 0 && math.Mod(float64(arg.Column+1), 2) == 0 {
				v[0], v[1], v[2] = byte(arg.Line*20), byte(arg.Column*100), byte(arg.Line+100)
				cd.Foreground(cd.EncodeColor(v[0], v[1], v[2]))
				cd.Box(int(arg.Xmax), int(arg.Xmin), int(arg.Ymax), int(arg.Ymin))
			} else {
				v[0], v[1], v[2] = cd.DecodeColor(cd.Background())
			}
			if scale >= 30 {
				cd.Foreground(cd.EncodeColor(best_contrast(v[0]), best_contrast(v[1]), best_contrast(v[2])))
				//cd.Box(int(arg.Xmax), int(arg.Xmin), int(arg.Ymax), int(arg.Ymin))
				cd.TextAlignment("CD_CENTER")
				cd.Font("Tahoma", "", 8)
				cd.Text(int(arg.Xmax+arg.Xmin)/2, int(arg.Ymax+arg.Ymin)/2, fmt.Sprintf("%d\n%d\n%d", v[0], v[1], v[2]))

			}
		},
		func(arg *iup.CellsNcols) {
			arg.Return = 4000
		},
		func(arg *iup.CellsNlines) {
			arg.Return = 1500
		},
		func(arg *iup.CellsHeight) {
			arg.Return = scale
		},
		func(arg *iup.CellsWidth) {
			arg.Return = scale
		},
		// func(arg *iup.CellsHspan) {
		// 	if arg.Line == 1 && arg.Column == 1 {
		// 		arg.Return = 2
		// 	} else if arg.Line == 5 && arg.Column == 5 {
		// 		arg.Return = 2
		// 	} else {
		// 		arg.Return = 1
		// 	}
		// },
		// func(arg *iup.CellsVspan) {
		// 	if arg.Line == 1 && arg.Column == 1 {
		// 		arg.Return = 2
		// 	} else if arg.Line == 5 && arg.Column == 5 {
		// 		arg.Return = 2
		// 	} else {
		// 		arg.Return = 1
		// 	}
		// },
	)

	res := iup.Vbox(
		"TABTITLE="+name,
		iup.Hbox(mtx, iup.Fill(),
			iup.Text("SPIN=YES,SPINVALUE=10,SPINMIN=1,SPINMAX=60,SPININC=2", func(arg *iup.TextSpin) {
				scale = arg.Inc
				fscale = true
			}),
		),
		cells,
	)
	iup.Timer("TIME=500,RUN=YES",
		func(arg *iup.TimerAction) {
			mtx.SetAttribute("1:2", mlc.Column)
			mtx.SetAttribute("1:5", mlc.Line)
			mtx.SetAttribute("1:3", float64(mlc.Column)/40)
			mtx.SetAttribute("1:6", float64(mlc.Line)/15)
			mtx.Redraw(1)
			if fscale {
				fscale = false
				cells.SetAttribute("REPAINT", "")
			}
			if fclick {
				fclick = false
				cells.SetAttribute("ORIGIN", fmt.Sprintf("%d:%d", mlc.Line, mlc.Column))
			}
		},
	)
	return res
}
func best_contrast(r byte) byte {
	if r < 128 {
		r = 255
	} else {
		r = 0
	}
	return r
}
