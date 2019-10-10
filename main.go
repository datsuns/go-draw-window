// ref) for windows API
//	https://mattn.kaoriya.net/software/lang/go/20160926174552.htm
// ref) for Desktop Drawing
//	https://stackoverflow.com/questions/1536141/how-to-draw-directly-on-the-windows-desktop-c
//	https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-beginpaint

package main

import (
	"fmt"
)

type HDC uintptr

type POINT struct {
	x int32
	y int32
}

type RECT struct {
	left   int32
	top    int32
	right  int32
	bottom int32
}

type PAINTSTRUCT struct {
	hdc         HDC
	fErase      bool
	rcPaint     RECT
	fRestore    bool
	fIncUpdate  bool
	rgbReserved [32]int8
}

func try_with_getdc() {
	prev := POINT{}
	r1, err := GetDC(0)
	fmt.Println("GetDC() -> ", r1)
	fmt.Println(err)
	MoveToEx(r1, 200, 200, &prev)
	r_, err := LineTo(r1, 500, 500)
	fmt.Println("LineTo() -> ", r_)
	r2, err := ReleaseDC(0, r1)
	fmt.Println("ReleaseDC() -> ", r2)
	fmt.Println(err)
}

func try_direct_attack() {
	prev := POINT{}
	p := PAINTSTRUCT{}
	hwnd, err := GetDesktopWindow()
	fmt.Println("GetDesktopWindow() -> ", hwnd, " ", err)
	hdc, err := BeginPaint(hwnd, &p)
	fmt.Println("BeginPaint() -> ", hdc, " ", err)

	MoveToEx(hdc, 20, 20, &prev)
	LineTo(hdc, 50, 50)
	MoveToEx(hdc, prev.x, prev.y, &prev)

	ok, err := EndPaint(hwnd, &p)
	fmt.Println("EndPaint() -> ", ok, " ", err)
}

func main() {
	try_direct_attack()
}
