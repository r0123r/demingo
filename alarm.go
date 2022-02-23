package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/r0123r/go-iup/iup"
)

func vbox9(name string) *iup.Handle {
	h := iup.Hbox(
		"TABTITLE="+name,
		"MARGIN=5x5,GAP=5",
		iup.BackgroundBox(
			"BGCOLOR=\"220 220 0\",BORDER=Yes",
			iup.Vbox(
				iup.Label(
					"EXPAND=HORIZONTAL,ALIGNMENT=ACENTER,FONTSIZE=12",
					iup.Attr("TITLE", "Welcome to GO-IUP"),
				),
				iup.Label(
					"EXPAND=YES,ALIGNMENT=ACENTER:ACENTER,WORDWRAP=YES",
					iup.Attr("TITLE",
						fmt.Sprintf("%s\nVersion: %s\n\n%s\nVersion: %s",
							iup.IupName, iup.IupVersion, iup.Name, iup.Version)),
				),
			),
		),
		iup.Vbox(
			iup.Button(
				"TITLE=Exit",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					arg.Return = iup.CLOSE
				},
			),
			iup.Button(
				"TITLE=About",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					iup.Message("About", "GO-IUP\nvisualfc@gmail.com 2011-2012")
				},
			),
			iup.Button(
				"TITLE=Message2",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					res := iup.Alarm("About", "GO-IUP\nvisualfc@gmail.com 2011-2012", "Да\n(Enter)", "Отмена\n(Esc)")
					iup.Message("Итог", fmt.Sprintf("Нажата кнопка %d", res))
				},
			),
			iup.Button(
				"TITLE=Message3",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					res := iup.Alarm("About", "GO-IUP\nvisualfc@gmail.com 2011-2012", "Да", "Нет", "Отмена")
					iup.Message("Итог", fmt.Sprintf("Нажата кнопка %d", res))
				},
			),
		),
		iup.Vbox(
			iup.Button(
				"TITLE=Error",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					m := iup.MessageDlg(`BUTTONS=OK,DIALOGTYPE=ERROR,TITLE=Error,VALUE="GO-IUP\nvisualfc@gmail.com 2011-2012"`)
					m.PopupA()
				},
			),
			iup.Button(
				"TITLE=INFORMATION",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					m := iup.MessageDlg(`BUTTONS=OK,DIALOGTYPE=INFORMATION,TITLE=Error,VALUE="GO-IUP\nvisualfc@gmail.com 2011-2012"`)
					m.PopupA()
				},
			),
			iup.Button(
				"TITLE=WARNING",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					m := iup.MessageDlg(`BUTTONS=OK,DIALOGTYPE=WARNING,TITLE=Error,VALUE="GO-IUP\nvisualfc@gmail.com 2011-2012"`)
					m.PopupA()
				},
			),
			iup.Button(
				"TITLE=QUESTION",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					m := iup.MessageDlg(`BUTTONS=OK,DIALOGTYPE=QUESTION,TITLE=Error,VALUE="GO-IUP\nvisualfc@gmail.com 2011-2012"`)
					m.PopupA()
				},
			),
		),
		iup.Vbox(
			iup.Button(
				"TITLE=GetParam",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					m := [][]string{
						{"String: %s\n", "true"},
						{"Multiline: %m\n", "second text\nsecond line"},
						{"Sep1 %t\n", ""},
						{"Color: %c{Color Tip}\n", "255 0 128"},
						{"Шрифт: %n\n", "Courier, 24"},
						{"File: %f[OPEN|*.bmp;*.jpg|CURRENT|NO|NO]\n", ""},
						{"Sep2 %t\n", ""},
						{"Options: %o|item0|item1|item2|\n", ""},
						{"List: %l|item0|item1|item2|item3|item4|item5|item6|\n", ""},
					}
					res := iup.GetParam("Задать параметры",
						strings.Join([]string{
							"%u[Да,No,Помощь]\n",
							m[0][0], m[1][0], m[2][0], m[3][0], m[4][0],
							m[5][0], //m[6][0], m[7][0], m[8][0], m[9][0],
						}, ""),
						&m[0][1], &m[1][1] /*&m[2][1],*/, &m[3][1], &m[4][1], &m[5][1])
					if res {
						fmt.Printf("%+v", m)
					}

				},
			),
			iup.Button(
				"TITLE=GetText",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					res, ok := iup.GetText("Title", "Введите текст ниже:\n")
					fmt.Println(ok, res)

				},
			),
			iup.Button(
				"TITLE=Custom",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					ok, res := GetTextParam("title", "value", "OK", "Cancel")
					fmt.Println(ok, res)

				},
			),
			iup.Button(
				"TITLE=ListDialog",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {

					marks := []int32{0, 0, 0, 0, 1, 1, 0, 0}
					options := []string{
						"Blue",
						"Red",
						"Green",
						"Yellow",
						"Black",
						"White",
						"Gray",
						"Brown"}

					error := iup.ListDialog(2, "Color Selection", options, 0, 16, 5, marks)

					if error == -1 {
						iup.Message("IupListDialog", "Operation canceled")
					} else {
						s := make([]string, 0)

						for i, m := range marks {
							if m > 0 {
								s = append(s, options[i])
							}
						}

						if len(s) > 0 {
							iup.Message("Options selected", strings.Join(s, "\n"))
						} else {
							iup.Message("IupListDialog", "No option selected")
						}
					}

				},
			),
			iup.Button(
				"TITLE=ColorDlg",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					h := iup.ColorDlg()
					h.PopupA()
				},
			),
			iup.Button(
				"TITLE=ProgressDlg",
				"SIZE=50x",
				func(arg *iup.ButtonAction) {
					log.Println("TITLE=ProgressDlg")
					dlg := iup.ProgressDlg()
					// dlg.SetCallback(func(arg *iup.CommonCancel) {
					// 	fmt.Print("CANCEL")
					// 	fmt.Println("Cancel")
					// 	arg.Return = iup.CONTINUE
					// })

					t := iup.Timer(func(arg *iup.TimerAction) {

						dlg.SetAttribute("INC", "")
						fmt.Println("timer", dlg.GetAttribute("COUNT"), dlg.GetAttribute("TOTALCOUNT"))
						if dlg.GetAttribute("COUNT") == dlg.GetAttribute("TOTALCOUNT") {
							dlg.Hide()
							arg.Return = iup.CLOSE
						}

						arg.Return = iup.CONTINUE

					})
					t.SetAttribute("TIME", "100")
					dlg.SetAttribute("TITLE", "IupProgressDlg Test")
					dlg.SetAttribute("DESCRIPTION", "Description first line\nSecond Line") // Actually can have any number of lines.
					//dlg.SetCallbackProc("CANCEL_CB", C.cancel_cb)                          //func() {
					// 	fmt.Print("CANCEL")
					// })
					dlg.SetAttribute("TOTALCOUNT", "30")
					t.SetAttribute("RUN", "YES")
					dlg.PopupA()
					t.Destroy()
				},
			),
		),
	)
	return h
}

func GetTextParam(title, value, title_ok, title_cancel string) (bool, string) {
	text := iup.Text(
		iup.Attrs("MULTILINE", "YES",
			"EXPAND", "YES",
			"VALUE", value,
			"FONT", "Courier,12",
			"VISIBLELINES", "10",
			"VISIBLECOLUMNS", "50",
		),
	)

	ok := iup.Button(
		iup.Attr("PADDING", "20x5"),
		iup.Attr("TITLE", title_ok),
		func(arg *iup.ButtonAction) {
			arg.Sender.GetDialog().(*iup.Handle).SetAttribute("STATUS", "1")
			arg.Return = iup.CLOSE
		},
	)

	cancel := iup.Button(
		iup.Attr("PADDING", "20x5"),
		iup.Attr("TITLE", title_cancel),
		func(arg *iup.ButtonAction) {
			arg.Sender.GetDialog().(*iup.Handle).SetAttribute("STATUS", "-1")
			arg.Return = iup.CLOSE
		},
	)

	dlg := iup.Dialog(
		"MINBOX=NO,MAXBOX=NO,SIZE=200x150",
		iup.Attr("TITLE", title),
		iup.Attr("ICON", iup.GetGlobal("ICON")),
		iup.Attr("PARENTDIALOG", iup.GetGlobal("PARENTDIALOG")),
		iup.Vbox(
			"MARGIN=10x10",
			"GAP=10",
			text,
			iup.Hbox(
				"MARGIN=0x0",
				"NORMALIZESIZE=HORZONTAL",
				iup.Fill(),
				ok,
				cancel,
			),
		),
	)
	defer dlg.Destroy()
	dlg.SetAttributeHandle("DEFAULTENTER", ok)
	dlg.SetAttributeHandle("DEFAULTESC", cancel)

	dlg.Map()

	text.SetAttribute("VISIBLELINES", "")
	text.SetAttribute("VISIBLECOLUMNS", "")

	dlg.Popup(iup.CENTERPARENT, iup.CENTERPARENT)
	if dlg.GetAttribute("STATUS") == "-1" {
		return false, ""
	}
	return true, text.GetAttribute("VALUE")
}
