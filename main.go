package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fnc := flag.Bool("funcs", false, "show built-in functions")
	ver := flag.Bool("version", false, "show version info")
	flag.Usage = func() {
		fmt.Printf("Equation Solver and Plotter %s\n", verinfo())
		fmt.Printf("\nUSAGE: %s [OPTIONS] [recipe.eqn]\n", filepath.Base(os.Args[0]))
		fmt.Println("\nOPTIONS")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *ver {
		fmt.Println(verinfo())
		os.Exit(0)
	}
	if *fnc {
		var fs []string
		for f := range lib {
			fs = append(fs, f)
		}
		sort.Slice(fs, func(i, j int) bool {
			return strings.ToLower(fs[i]) < strings.ToLower(fs[j])
		})
		fmt.Println(strings.Join(fs, ", "))
		os.Exit(0)
	}
	if len(flag.Args()) == 0 {
		fmt.Println("ERROR: missing recipe (-help for usage)")
		os.Exit(1)
	}
	recipe := flag.Arg(0)
	lf, err := os.Create(recipe + ".log")
	assert(err)
	defer func() {
		err := allege(recover())
		if err != nil {
			fmt.Fprintln(lf, err.Error())
		}
		assert(lf.Close())
		if err == nil {
			os.Remove(recipe + ".log")
		}
	}()
	r, err := LoadRecipe(recipe)
	assert(err)
	assert(r.Calculate())
	assert(r.Solve())
}
