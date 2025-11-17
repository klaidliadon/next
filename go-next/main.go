// Command line utility to generate combinations and permutations using next library.
package main

import (
	"flag"
	"fmt"
	"iter"
	"os"

	"klaidliadon.dev/next"
)

func main() {
	order := flag.Bool("order", false, "order matters")
	repeat := flag.Bool("repeat", false, "elements can be repeated")
	size := flag.Int("size", 0, "size of each result (not 0)")
	flag.Parse()

	if *size == 0 {
		flag.Usage()
		os.Exit(1)
	}

	base := flag.Args()

	if len(base) == 0 {
		fmt.Printf("%s: please specify at least an element.\n", os.Args[0])
		os.Exit(1)
	}

	seq := (func([]string, int) iter.Seq[[]string])(nil)
	switch {
	case *repeat && *order:
		seq = next.RepeatPermutation
	case *repeat && !*order:
		seq = next.RepeatCombination
	case !*repeat && *order:
		seq = next.Permutation
	case !*repeat && !*order:
		seq = next.Combination
	}

	for c := range seq(base, *size) {
		fmt.Println(c)
	}
}
