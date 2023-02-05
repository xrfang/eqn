package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ver := flag.Bool("version", false, "show version info")
	flag.Usage = func() {
		fmt.Printf("Equation Solver and Plotter %s\n", verinfo())
		fmt.Printf("\nUSAGE: %s [OPTIONS] <equation-file.eqn>\n", filepath.Base(os.Args[0]))
		fmt.Println("\nOPTIONS")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *ver {
		fmt.Println(verinfo())
		os.Exit(0)
	}
	if len(flag.Args()) == 0 {
		fmt.Println("ERROR: missing equation file (-help for usage)")
		os.Exit(1)
	}
	r, err := LoadRecipe(flag.Arg(0))
	assert(err)
	assert(r.Calculate())
	assert(r.Solve())
}
