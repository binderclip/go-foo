package main

import (
	"fmt"

	"github.com/binderclip/go-bar/bar"
	"github.com/binderclip/go-baz/baz"
)

func main() {
	fmt.Println("foo 001 >>")
	bar.Bar()
	baz.Baz()
	fmt.Println("foo 001 <<")
}
