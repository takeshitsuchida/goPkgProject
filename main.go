package main

import (
	"fmt"
	"goPkgProject/mypkg"
	"goPkgProject/mypkg/underpkg"
)

func main() {
	fmt.Println(mypkg.Rectangle(4, 5))
	mypkg.Intro()

	underpkg.Hello()
}
