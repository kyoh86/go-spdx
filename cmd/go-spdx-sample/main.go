// +build sample

package main

import (
	"fmt"

	"github.com/kyoh86/go-spdx/spdx"
)

func main() {
	tree, err := spdx.Parse("0BSD AND (0BSD OR 0BSD)")
	if err != nil {
		panic(err)
	}
	fmt.Println(tree)

	mit, err := spdx.Get("MIT")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", *mit)
}
