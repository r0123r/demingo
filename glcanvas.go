package main

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/r0123r/go-iup/iup"
	"github.com/r0123r/go-iup/iupgl"
)

var t float32
var finit bool

func vbox72(name string) *iup.Handle {
	iupgl.CanvasOpen()
	cnv := iupgl.Canvas(
		"RASTERSIZE=640x480", "DEPTH_SIZE=16",
		"BUFFER=DOUBLE",
		func(arg *iup.CanvasAction) {
			iupgl.MakeCurrent(arg.Sender)
			//gl.Init()
			gl.ClearColor(0.3, 0.3, 0.3, 1.0) /* White */
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			gl.Enable(gl.DEPTH_TEST)

			gl.MatrixMode(gl.MODELVIEW)
			gl.PushMatrix() /* saves current model view in a stack */
			gl.Translatef(0.0, 0.0, 0.0)
			gl.Scalef(1.0, 1.0, 1.0)
			gl.Rotatef(t, 0, 0, 1)
			colorCube()
			gl.PopMatrix()

			iupgl.SwapBuffers(arg.Sender)
		},
		func(arg *iup.CanvasResize) {
			iupgl.MakeCurrent(arg.Sender) /* Make the canvas current in OpenGL */
			if !finit {
				finit = true
				gl.Init()
			}
			/* define the entire canvas as the viewport  */
			gl.Viewport(0, 0, arg.Width, arg.Height)

			/* transformation applied to each vertex */
			gl.MatrixMode(gl.MODELVIEW)
			gl.LoadIdentity() /* identity, i. e. no transformation */

			/* projection transformation (orthographic in the xy plane) */
			gl.MatrixMode(gl.PROJECTION)
			gl.LoadIdentity()
			iupgl.Perspective(60, 4/3, 1, 15)
			iupgl.LookAt(3, 3, 3, 0, 0, 0, 0, 0, 1)

		},
		func(arg *iup.CommonDestroy) {
			fmt.Println("Exit cube")
		},
	)
	iup.Timer("TIME=50,RUN=YES",
		func(arg *iup.TimerAction) {
			t++
			cnv.Redraw(1)
		})
	res := iup.Vbox(
		"TABTITLE="+name,
		cnv,
	)
	return res
}
func polygon(a, b, c, d int) {
	vertices := [][3]float64{
		{-1, -1, 1},
		{-1, 1, 1},
		{1, 1, 1},
		{1, -1, 1},
		{-1, -1, -1},
		{-1, 1, -1},
		{1, 1, -1},
		{1, -1, -1}}
	gl.Begin(gl.POLYGON)
	gl.Vertex3dv(&vertices[a][0])
	gl.Vertex3dv(&vertices[b][0])
	gl.Vertex3dv(&vertices[c][0])
	gl.Vertex3dv(&vertices[d][0])
	gl.End()
}

func colorCube() {
	gl.Color3f(1, 0, 0)
	gl.Normal3f(1, 0, 0)
	polygon(2, 3, 7, 6)

	gl.Color3f(0, 1, 0)
	gl.Normal3f(0, 1, 0)
	polygon(1, 2, 6, 5)

	gl.Color3f(0, 0, 1)
	gl.Normal3f(0, 0, 1)
	polygon(0, 3, 2, 1)

	gl.Color3f(1, 0, 1)
	gl.Normal3f(0, -1, 0)
	polygon(3, 0, 4, 7)

	gl.Color3f(1, 1, 0)
	gl.Normal3f(0, 0, -1)
	polygon(4, 5, 6, 7)

	gl.Color3f(0, 1, 1)
	gl.Normal3f(-1, 0, 0)
	polygon(5, 4, 0, 1)
}
