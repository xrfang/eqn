package main

import "github.com/Knetic/govaluate"

type (
	Expression struct {
		n string
		v *govaluate.EvaluableExpression
	}
	Function struct {
		IVar string   `yaml:"ivar"`
		Expr []string `yaml:"expr"`
		expr []Expression
	}
)
