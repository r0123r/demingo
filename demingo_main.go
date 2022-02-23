// Copyright (C) 2011-2012 r0123r. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the COPYRIGHT file.

package main

import (
	"fmt"

	"github.com/r0123r/go-iup/iup"
	"github.com/r0123r/go-iup/iup_plot"
	"github.com/r0123r/go-iup/iupctls"
	"github.com/r0123r/go-iup/iupimglib"
)

func main() {
	e := iup.Open()
	if e != nil {
		fmt.Println(e)
		return
	}
	defer iup.Close()

	iupctls.Open()
	iup_plot.Open()
	iupimglib.Open()

	ui()
}

func ui() {

	tree := iup.Tree(
		"EXPANDALL=YES",
	)
	tree.SetName("tree")
	pr := iup.ProgressBar("MIN=0,MAX=100,VALUE=50,EXPAND=HORIZONTAL")
	vbox1 := iup.Vbox(
		"TABTITLE=Standard,MARGIN=0x0,GAP=5",
		iup.Split("ORIENTATION=VERTICAL,VALUE=500",
			iup.Frame("TITLE=Frame,EXPAND=YES",
				iup.Vbox(
					iup.Label("TITLE=\"List Label\""),
					iup.List("EXPAND=YES,1=list1", "2=list2", "3=list3"),
				),
			),
			iup.Frame(
				iup.Radio(
					iup.Hbox(
						iup.Toggle("TITLE=1"),
						iup.Toggle("TITLE=2"),
						iup.Toggle("TITLE=3"),
						iup.Toggle("TITLE=4"),
					),
				),
			),
		),
		iup.Hbox(
			iup.List("DROPDOWN=YES,EDITBOX=YES", "1=list1", "2=list2"),
			iup.SpinBox(
				iup.Text("VALUE=100,SPIN=YES"),
			),
			iup.Toggle("TITLE=Toggle,EXPAND=VERTICAL"),
			iup.Val("MIN=1,MAX=100"),
		),
		iup.Text("EXPAND=HORIZONTAL,SPIN=YES,SPINVALUE=50", func(arg *iup.TextSpin) {
			pr.SetAttribute("VALUE", fmt.Sprint(arg.Inc))
			arg.Return = iup.DEFAULT
		}),
		pr,
		iup.Sbox("DIRECTION=NORTH,SIZE=0x800",
			tree,
		),
	)
	vbox2 := iup.Vbox(
		"TABTITLE=Text",
		iup.Text("EXPAND=YES,MULTILINE=YES"),
	)
	vbox3 := iup.Vbox(
		"TABTITLE=Canvas",
		iup.Canvas(
			func(arg *iup.CanvasAction) {
				d := iup.NewDraw(arg.Sender)
				var w, h int

				d.Begin()

				d.GetSize(&w, &h)

				/* white background */
				d.SetAttribute("DRAWCOLOR", "255 255 255")
				//  IupSetAttribute(ih, "DRAWCOLOR", "255 0 255"); /* pink */
				d.SetAttribute("DRAWSTYLE", "FILL")
				d.Rectangle(0, 0, w-1, h-1)

				/* Guide Lines */
				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.SetAttribute("DRAWSTYLE", "STROKE")
				d.Line(10, 5, 10, 19)
				d.Line(14, 5, 14, 19)
				d.Line(5, 10, 19, 10)
				d.Line(5, 14, 19, 14)

				/* Stroke Rectangle, must cover guide lines */
				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.SetAttribute("DRAWSTYLE", "STROKE")
				d.Rectangle(10, 10, 14, 14)

				// /* Guide Lines */
				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.Line(10, 5+30, 10, 19+30)
				d.Line(14, 5+30, 14, 19+30)
				d.Line(5, 10+30, 19, 10+30)
				d.Line(5, 14+30, 19, 14+30)

				// /* Fill Rectangle, must cover guide lines */
				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.SetAttribute("DRAWSTYLE", "FILL")
				d.Rectangle(10, 10+30, 14, 14+30)

				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.Rectangle(30, 10, 50, 30)

				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.Arc(30, 10, 50, 30, 0, 360)

				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.Rectangle(60, 10, 80, 30)

				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.SetAttribute("DRAWSTYLE", "FILL")
				d.Arc(60, 10, 80, 30, 0, 360)

				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.Rectangle(30, 10+30, 50, 30+30)

				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.Arc(30, 10+30, 50, 30+30, 45, 135)

				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.Rectangle(60, 10+30, 80, 30+30)

				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.SetAttribute("DRAWSTYLE", "FILL")
				d.Arc(60, 10+30, 80, 30+30, 45, 135)

				d.SetAttribute("DRAWCOLOR", "255 0 0")
				d.Line(20, 70-2, 20, 70+2)
				d.Line(20-2, 70, 20+2, 70)

				d.SetAttribute("DRAWCOLOR", "0 0 0")
				d.SetAttribute("DRAWFONT", "Helvetica, -30")
				d.SetAttribute("DRAWTEXTORIENTATION", "60")
				d.SetAttribute("DRAWTEXTLAYOUTCENTER", "Yes")
				d.GetTextSize("Text", &w, &h)
				d.SetAttribute("DRAWSTYLE", "STROKE")
				d.Rectangle(20, 70, 20+w, 70+h)
				d.Text("Text", 20, 70, -1, -1)
				d.SetAttribute("DRAWTEXTORIENTATION", "0")
				d.SetAttribute("DRAWTEXTWRAP", "Yes")
				d.SetAttribute("DRAWTEXTELLIPSIS", "Yes")
				d.SetAttribute("DRAWTEXTCLIP", "Yes")
				d.Text("Very Large Text", 40, 190, 50, 70)

				// //  IupSetAttribute(ih, "DRAWLINEWIDTH", "3");
				d.SetAttribute("DRAWSTYLE", "STROKE")
				d.Line(10, 110, 100, 110)
				d.SetAttribute("DRAWSTYLE", "STROKE_DASH")
				d.Line(10, 110+5, 100, 110+5)
				d.SetAttribute("DRAWSTYLE", "STROKE_DOT")
				d.Line(10, 110+10, 100, 110+10)
				d.SetAttribute("DRAWSTYLE", "STROKE_DASH_DOT")
				d.Line(10, 110+15, 100, 110+15)
				d.SetAttribute("DRAWSTYLE", "STROKE_DASH_DOT_DOT")
				d.Line(10, 110+20, 100, 110+20)

				d.Image("IUP_FileSave", 110, 10, -1, -1)
				d.Image("IUP_FileSave", 110, 40, 50, 50)
				//d.Image("Test32bpp", 110, 70, -1, -1)
				// //  IupDrawImage(ih, "Test32bpp", 110, 70, 60, 60);

				d.SetAttribute("DRAWFONT", "Helvetica, Bold -15")
				d.Text(d.GetAttribute("DRAWDRIVER"), 70, 135, -1, -1)

				d.End()

			},
		),
	)
	vbox4 := iup.Vbox(
		"TABTITLE=ColorBrowser",
		iup.ColorBrowser(
			"EXPAND=YES",
			func(arg *iup.ColorBrowserDrag) {
				fmt.Printf("R=%d,G=%d,B=%d\n", arg.R, arg.G, arg.B)
			},
		),
		iup.Frame(`TITLE="Выбор цвета в палитре",EXPAND=YES`,
			iup.Colorbar(
				"ORIENTATION=HORIZONTAL",
				"NUM_PARTS=2",
				"SHOW_SECONDARY=YES",
				func(arg *iup.ColorbarSelect) {
					fmt.Println(arg)
				},
			),
		),
	)
	file_menu := iup.Menu(
		iup.Item("Open"),
		iup.Item("SaveAs"),
		iup.Separator(),
		iup.Item("Exit", func(arg *iup.ItemAction) {

		}),
	)

	dlg := iup.Dialog(
		iup.Attr("TITLE", "GO-IUP Demo 1.0"),
		"SIZE=350x200",
		"SHRINK=YES",
		"MARGIN=10x10",
		"PLACEMENT=MAXIMIZED",
		//"FULLSCREEN=YES",
		iup.Tabs(
			vbox1,
			vbox2,
			vbox3,
			vbox4,
			vbox5("Dial"),
			vbox6("Plot"),
			vbox7("Matrix"),
			vbox71("Cells"),
			vbox72("GLcanvas"),
			vbox8("Image"),
			vbox9("Messagebox"),
		),
	)
	dlg.SetAttributeHandle("MENU", file_menu)
	dlg.Show()
	tree.SetAttribute("ADDBRANCH", "Item2")
	tree.SetAttribute("ADDLEAF1", "leaf3")
	tree.SetAttribute("ADDLEAF2", "leaf4")
	tree.SetAttribute("ADDBRANCH", "Item1")
	tree.SetAttribute("ADDLEAF1", "leaf1")
	tree.SetAttribute("ADDLEAF2", "leaf2")

	defer dlg.Destroy()
	iup.MainLoop()
}
