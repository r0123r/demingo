package main

import (
	"fmt"
	"math"

	"github.com/r0123r/go-iup/iup"
	"github.com/r0123r/go-iup/iup_plot"
)

func vbox6(name string) *iup.Handle {
	tt := 6500
	xx := make([]float64, tt)
	y0 := make([]float64, tt)
	y1 := make([]float64, tt)
	y2 := make([]float64, tt)
	lx := iup.Text("READONLY=YES,BORDER=NO,SIZE=50x10")
	ly0 := iup.Text("READONLY=YES,BORDER=NO,SIZE=50x10,VISIBLE=NO")
	ly1 := iup.Text("READONLY=YES,BORDER=NO,SIZE=50x10,VISIBLE=NO")
	ly2 := iup.Text("READONLY=YES,BORDER=NO,SIZE=50x10,VISIBLE=NO")
	plot := iup_plot.PPlot(
		iup.Attrs(
			"MARGINBOTTOM", "35",
			"MARGINLEFT", "35",
			"SHOWCROSSHAIR", "HORIZONTAL",
			"LEGENDSHOW", "NO",
			"LEGENDPOS", "BOTTOMCENTER",
			"LEGENDBOX", "NO",
			"GRID", "YES",
			"MENUITEMPROPERTIES", "YES",
			"BOX", "YES",
		),
		func(arg *iup.PlotMotion) {
			lx.SetAttribute("VALUE", fmt.Sprintf("%e", arg.X))
			for i, xi := range xx {
				if xi >= arg.X {
					ly0.SetAttribute("VALUE", fmt.Sprintf("%e", y0[i]))
					ly1.SetAttribute("VALUE", fmt.Sprintf("%e", y1[i]))
					ly2.SetAttribute("VALUE", fmt.Sprintf("%e", y2[i]))
					break
				}
			}

		},
	)
	c := 0
	for x := -3.14; x <= 3.14; x += .001 {
		xx[c] = x
		y0[c] = math.Sin(x)
		y1[c] = math.Sin(x) * math.Cosh(x)
		y2[c] = math.Sin(x) * math.Sinh(x)
		c++
	}
	index := 0
	plot_count := 1
	res := iup.Vbox(
		"TABTITLE="+name,
		iup.Hbox(
			iup.Button("TITLE=Exit", "SIZE=50x",
				func(arg *iup.ButtonAction) {
					arg.Return = iup.CLOSE
				},
			),
			iup.Button("TITLE=AddPlot", "SIZE=50x",
				func(arg *iup.ButtonAction) {
					plot_count++
					plot.SetAttribute("PLOT_COUNT", "2")
					plot.SetAttribute("PLOT_CURRENT", "1")
					plot.SetAttribute("BOX", "YES")
					plot.SetAttribute("REDRAW", "")
				},
			),
			iup.Label("TITLE=X:,SIZE=20x10,ALIGNMENT=ACENTER"), lx,
			iup.Toggle(`TITLE="plot0",SIZE=30x10,ALIGNMENT=ACENTER:ARIGHT`,
				func(arg *iup.ToggleAction) {
					if arg.State > 0 {
						plot.Begin(0)
						plot.End()
						plot.PlotAddSamples(index, &xx[0], &y0[0], c)
						ly0.SetAttribute("VISIBLE", "YES")
						plot.SetAttribute("DS_NAME", "plot0")
						index++
					} else {
						plot.SetAttribute("REMOVE", "plot0")
						ly0.SetAttribute("VISIBLE", "NO")
						index--
					}
					plot.SetAttribute("DS_ORDEREDX", "YES")
					plot.SetAttribute("REDRAW", "")
				},
			), ly0,
			iup.Toggle(`TITLE="plot1",SIZE=30x10,ALIGNMENT=ACENTER:ARIGHT`,
				func(arg *iup.ToggleAction) {
					if arg.State > 0 {
						plot.Begin(0)
						plot.End()
						plot.PlotAddSamples(index, &xx[0], &y1[0], c)
						index++
						plot.SetAttribute("DS_NAME", "plot1")
						ly1.SetAttribute("VISIBLE", "YES")
					} else {
						plot.SetAttribute("REMOVE", "plot1")
						ly1.SetAttribute("VISIBLE", "NO")
						index--
					}
					plot.SetAttribute("DS_ORDEREDX", "YES")
					plot.SetAttribute("REDRAW", "")
				},
			), ly1,
			iup.Toggle(`TITLE="plot2",SIZE=30x10,ALIGNMENT=ACENTER:ARIGHT`,
				func(arg *iup.ToggleAction) {
					if arg.State > 0 {
						plot.Begin(0)
						plot.End()
						plot.PlotAddSamples(index, &xx[0], &y2[0], c)
						index++
						plot.SetAttribute("DS_NAME", "plot2")
						ly2.SetAttribute("VISIBLE", "YES")
					} else {
						plot.SetAttribute("REMOVE", "plot2")
						ly2.SetAttribute("VISIBLE", "NO")
						index--
					}
					plot.SetAttribute("DS_ORDEREDX", "YES")
					plot.SetAttribute("REDRAW", "")
				},
			), ly2,
		),
		plot,
	)
	return res
}
