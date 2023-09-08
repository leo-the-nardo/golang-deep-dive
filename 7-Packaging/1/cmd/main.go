package main

import (
	"fmt"
	"github.com/leo-the-nardo/goexpert/7-Packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}
