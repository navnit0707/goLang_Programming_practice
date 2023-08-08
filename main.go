package main

import (
	"fmt"
	"os"
)

func main(){

	var s,sep string
	for i:= 0;i< len(os.Args);i++{
		s += sep+ os.Args[i]
		sep = " "

	}
	fmt.Println(s)
	// result will be :   C:\Users\navni\AppData\Local\Temp\go-build319144518\b001\exe\main.exe hello hi
	// Arg[0] is the name of command itself
}