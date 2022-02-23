package main

import (
	"github.com/r0123r/go-iup/iup"
)

func vbox8(name string) *iup.Handle {
	var img *iup.Handle
	res := iup.Vbox(
		"TABTITLE="+name,
		iup.Hbox(
			"GAP=10",
			iup.Button("IMAGE=IUP_FileNew"),
			iup.Button("IMAGE=IUP_FileOpen"),
			iup.Button(
				"TITLE=SelectImg",
				func(arg *iup.ButtonAction) {
					file, ok := iup.GetOpenFile("", "*.png;*.jpg;*.bmp;*.jpeg")
					if ok {
						label := arg.Sender.GetDialogChild("img_label")
						if img != nil {
							img.Destroy()
						}
						img = iup.LoadImage(file)
						if img != nil {
							img.SetName("img_label_image")
							label.(*iup.Handle).SetAttribute("IMAGE", "img_label_image")
							label.(*iup.Handle).Refresh()
						}
					}
				},
			),
		),
		iup.Label("IMAGE=img_label_image", "NAME=img_label"),
	)
	return res
}
