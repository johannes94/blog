package main

import (
	"fmt"
	"init-functions/b"
	"init-functions/a"
)

func init() {
	fmt.Println("pkg main file init")
}

func main() {
	fmt.Println("func main()")
	fmt.Println(a.A)
	fmt.Println(b.B)
}
