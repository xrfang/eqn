package main

import (
	"math"

	"github.com/Knetic/govaluate"
)

var lib map[string]govaluate.ExpressionFunction

func init() {
	lib = map[string]govaluate.ExpressionFunction{
		"sin": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Sin(x), nil
		},
		"cos": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Cos(x), nil
		},
		"tan": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Tan(x), nil
		},
		"sinh": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Sinh(x), nil
		},
		"cosh": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Cosh(x), nil
		},
		"tanh": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Tanh(x), nil
		},
		"asin": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Asin(x), nil
		},
		"acos": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Acos(x), nil
		},
		"atan": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Atan(x), nil
		},
		"asinh": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Asinh(x), nil
		},
		"acosh": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Acosh(x), nil
		},
		"atanh": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Atanh(x), nil
		},
		"erf": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Erf(x), nil
		},
		"erfi": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Erfinv(x), nil
		},
		"erfc": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Erfc(x), nil
		},
		"erfci": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Erfcinv(x), nil
		},
		"exp": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Exp(x), nil
		},
		"gamma": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Gamma(x), nil
		},
		"ln": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Log(x), nil
		},
		"log": func(args ...any) (any, error) {
			b := args[0].(float64)
			x := args[1].(float64)
			return math.Log(x) / math.Log(b), nil
		},
		"J0": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.J0(x), nil
		},
		"J1": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.J1(x), nil
		},
		"Jn": func(args ...any) (any, error) {
			n := int(args[0].(float64))
			x := args[1].(float64)
			return math.Jn(n, x), nil
		},
		"Y0": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Y0(x), nil
		},
		"Y1": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Y1(x), nil
		},
		"Yn": func(args ...any) (any, error) {
			n := int(args[0].(float64))
			x := args[1].(float64)
			return math.Yn(n, x), nil
		},
		"sqrt": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Sqrt(x), nil
		},
		"abs": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Abs(x), nil
		},
		"ceil": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Ceil(x), nil
		},
		"floor": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Floor(x), nil
		},
		"round": func(args ...any) (any, error) {
			x := args[0].(float64)
			return math.Round(x), nil
		},
	}
}
