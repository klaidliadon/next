// Command line utility to generate combinations and permutations using next library.
package main

import (
	"flag"
	"fmt"

	"github.com/klaidliadon/next"
)

func main() {
	var o bool
	var r bool
	var s uint = 3
	flag.BoolVar(&o, "order", false, "order matters")
	flag.BoolVar(&r, "repeat", false, "repeatition")
	flag.UintVar(&s, "size", 0, "size of each result")
	flag.Parse()
	var args []interface{}
	for _, v := range flag.Args() {
		args = append(args, v)
	}
	var ch <-chan []interface{}
	if o {
		ch = next.Permutation(args, int(s), r)
	} else {
		ch = next.Combination(args, int(s), r)
	}
	for c := range ch {
		fmt.Println(c)
	}
}
