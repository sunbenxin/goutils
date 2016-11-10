package main

import (
	//ch "github.com/goutils"
	"fmt"
	"github.com/goutils/test"
)

func check() {
	//set up the pipeline
	c := test.Gen(2, 3)
	out := test.Sq(c)
	//consume the output
	test.Consumer(out)
}

func main() {
	fmt.Println("dasf")
	test.Mouse()
}
