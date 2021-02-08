package main

import (
	"fmt"
	"github.com/kaustavc/test-go-project1/packagea/mathlib"
	"github.com/kaustavc/test-go-project1/packageb"
)

func main() {
	x := mathlib.Add(1, 2)
	s := packageb.ToCaps("hello")
	fmt.Println(x, s)
}
