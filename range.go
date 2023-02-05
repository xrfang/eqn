package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type (
	Range  []float64
	Ranges map[string]Range
)

func (r *Range) UnmarshalYAML(um func(interface{}) error) (err error) {
	defer func() { err = allege(recover()) }()
	var expr string
	assert(um(&expr))
	f := func(s string) float64 {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(fmt.Errorf("range: '%v'不是合法的浮点数", s))
		}
		return v
	}
	for _, x := range strings.Split(expr, ",") {
		m := xr.FindStringSubmatch(x)
		assert(len(m) > 0, fmt.Sprintf("'%s'不是合法的范围表达式", x))
		min := f(m[1])
		max := min
		if m[2] != "" {
			max = f(m[2])
		}
		step := float64(1)
		if m[3] != "" {
			step = f(m[3])
		}
		if step == 0 {
			*r = append(*r, min, max)
		} else {
			for {
				*r = append(*r, min)
				min += step
				if min >= max+step/2 {
					break
				}
			}
		}
	}
	sort.Slice(*r, func(i, j int) bool { return (*r)[i] < (*r)[j] })
	return nil
}

func (rs Ranges) Params() []map[string]any {
	var params []map[string]any
	for n, r := range rs {
		if len(params) == 0 {
			for _, v := range r {
				params = append(params, map[string]any{n: v})
			}
		} else {
			var params2 []map[string]any
			for _, p := range params {
				for _, v := range r {
					p2 := make(map[string]any)
					for k, v := range p {
						p2[k] = v
					}
					p2[n] = v
					params2 = append(params2, p2)
				}
			}
			params = params2
		}
	}
	return params
}

var xr *regexp.Regexp

func init() {
	rangeRx := `^\s*([0-9.-]+)(?:\s*~\s*([0-9.-]+))?(?:\s*\+\s*([0-9.]+))?\s*$`
	xr = regexp.MustCompile(rangeRx)
}
