package main

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func main() {
	libuser32 := windows.NewLazySystemDLL("user32.dll")
	fGetDC := libuser32.NewProc("GetDC")
	fReleaseDC := libuser32.NewProc("ReleaseDC")

	fmt.Println(fGetDC)
	fmt.Println(fReleaseDC)
	r1, r2, err := fGetDC.Call(0)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(err)
}
