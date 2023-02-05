package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Precision int     `yaml:"precision"`
		Epsilon   float64 `yaml:"epsilon"`
	}
	Recipe struct {
		Conf Config
		Func *Function    `yaml:"function"`
		Calc *Calculation `yaml:"calculate"`
		Root *Calculation `yaml:"solve"`
		path string
	}
)

func LoadRecipe(fn string) (r *Recipe, err error) {
	defer func() { err = allege(recover()) }()
	f, err := os.Open(fn)
	assert(err)
	defer f.Close()
	r = new(Recipe)
	r.path, _ = filepath.Abs(fn)
	r.path = filepath.Dir(r.path)
	assert(yaml.NewDecoder(f).Decode(r))
	if r.Conf.Precision < 0 || r.Conf.Precision > 9 {
		r.Conf.Precision = 0
	}
	assert(r.Func != nil, "函数未定义")
	xr := regexp.MustCompile(`(?i)^([a-z][a-z0-9]{0,8})\s*=\s*(.*?)\s*$`)
	for _, def := range r.Func.Expr {
		def = strings.TrimSpace(def)
		if def == "" {
			continue
		}
		m := xr.FindStringSubmatch(def)
		assert(len(m) == 3, "表达式必须包含等号，左侧为英文字母与数字构成的变量名，右侧为定义")
		lhs := strings.TrimSpace(m[1])
		rhs := strings.TrimSpace(m[2])
		assert(len(rhs) > 0, "表达式定义不能为空")
		eval, err := govaluate.NewEvaluableExpressionWithFunctions(rhs, lib)
		assert(err)
		r.Func.expr = append(r.Func.expr, Expression{lhs, eval})
	}
	if r.Calc != nil {
		r.Calc.Validate(r)
	}
	if r.Root != nil {
		r.Root.Validate(r)
	}
	return
}
