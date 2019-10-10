package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Println("vim-go")
	beepFunc := syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
	fmt.Println(beepFunc)
}
