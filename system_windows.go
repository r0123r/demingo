package main

import (
	"syscall"
	"unsafe"
)

var (
	user32                   = syscall.NewLazyDLL("user32.dll")
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	getModuleHandle          = kernel32.NewProc("GetModuleHandleW")
	getConsoleWindow         = kernel32.NewProc("GetConsoleWindow")
	getCurrentProcessId      = kernel32.NewProc("GetCurrentProcessId")
	getWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
	showWindowAsync          = user32.NewProc("ShowWindowAsync")
)

var Exec_path string

func init() {
	//func hideConsole() {
	console := GetConsoleWindow()
	if console == 0 {
		return // no console attached
	}
	// If this application is the process that created the console window, then
	// this program was not compiled with the -H=windowsgui flag and on start-up
	// it created a console along with the main application window. In this case
	// hide the console window.
	// See
	// http://stackoverflow.com/questions/9009333/how-to-check-if-the-program-is-run-from-a-console
	_, consoleProcID := GetWindowThreadProcessId(console)
	if GetCurrentProcessId() == consoleProcID {
		ShowWindowAsync(console, 0) //SW_HIDE
	}
}

func GetConsoleWindow() uintptr {
	ret, _, _ := getConsoleWindow.Call()
	return ret
}
func GetWindowThreadProcessId(hwnd uintptr) (uintptr, uint32) {
	var processId uint32
	ret, _, _ := getWindowThreadProcessId.Call(
		hwnd,
		uintptr(unsafe.Pointer(&processId)),
	)
	return ret, processId
}
func ShowWindowAsync(window, commandShow uintptr) bool {
	ret, _, _ := showWindowAsync.Call(window, commandShow)
	return ret != 0
}
func GetCurrentProcessId() uint32 {
	id, _, _ := getCurrentProcessId.Call()
	return uint32(id)
}
