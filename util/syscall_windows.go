package main

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go syscall_windows.go

//sys   GetDC(hwnd uintptr) (hdc uintptr, err error) = user32.GetDC
//sys   ReleaseDC(hwnd uintptr, hdc uintptr) (succeeded int, err error) = user32.ReleaseDC
