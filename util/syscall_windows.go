package main

import (
	"syscall"
)

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go syscall_windows.go

//sys   GetDC(hwnd syscall.Handle) (hdc syscall.Handle, err error) = user32.GetDC
//sys   ReleaseDC(hwnd syscall.Handle, hdc syscall.Handle) (succeeded int, err error) = user32.ReleaseDC
//sys   BeginPaint(hwnd syscall.Handle, lpPaint *PAINTSTRUCT) (hdc syscall.Handle, err error) = user32.BeginPaint
//sys   EndPaint(hwnd syscall.Handle, lpPaint *PAINTSTRUCT) (succeeded int, err error) = user32.EndPaint
//sys   GetDesktopWindow() (hwnd syscall.Handle ,err error) = user32.GetDesktopWindow
//sys   LineTo(hwnd syscall.Handle, x int32, y int32) (succeeded int, err error) = gdi32.LineTo
//sys   MoveToEx(hwnd syscall.Handle, x int32, y int32, lppt *POINT) (succeeded int, err error) = gdi32.MoveToEx
