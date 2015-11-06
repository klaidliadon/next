// Command line utility to generate combinations and permutations using next library.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/klaidliadon/next"
)

func main() {
	var o bool
	var r bool
	var s uint
	flag.BoolVar(&o, "order", false, "order matters")
	flag.BoolVar(&r, "repeat", false, "repeatition")
	flag.UintVar(&s, "size", 0, "size of each result (not 0)")
	flag.Parse()
	var args []interface{}
	for _, v := range flag.Args() {
		args = append(args, v)
	}
	if s == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if len(args) == 0 {
		fmt.Printf("%s: please specify at least an element.\n", os.Args[0])
		os.Exit(1)
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
