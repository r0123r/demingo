module demingo

go 1.17

replace github.com/r0123r/go-iup => ../../../go-iup

replace github.com/r0123r/glut => ../../../glut

require (
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6
	github.com/r0123r/go-iup v0.0.0-00010101000000-000000000000
)

require github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211213063430-748e38ca8aec // indirect
