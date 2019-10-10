package main

import (
	"fmt"
)

func main() {
	r1, err := GetDC(0)
	fmt.Println(r1)
	fmt.Println(err)
	r2, err := ReleaseDC(0, r1)
	fmt.Println(r2)
	fmt.Println(err)
}
