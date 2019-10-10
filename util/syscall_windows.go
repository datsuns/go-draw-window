package main

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go syscall_windows.go

//sys   FormatMessage(flags uint32, source syscall.Handle, messageID uint32, languageID uint32, buffer *byte, bufferSize uint32, arguments uintptr) (numChars uint32, err error) = kernel32.FormatMessageW
