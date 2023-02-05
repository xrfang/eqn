package main

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
)

func (c *Calculation) bisect(iv string, param *map[string]any) {
	defer func() {
		switch r := recover().(type) {
		case float64:
			(*param)[iv] = r
		case error:
			panic(r)
		}
	}()
	calc := func(x float64) (y float64) {
		(*param)[iv] = x
		y, _ = c.call(param)
		if math.Abs(y) <= c.err {
			panic(x)
		}
		return y
	}
	minx := c.min
	maxx := c.max
	ss := func(v1, v2 float64) bool {
		return (v1 < 0 && v2 < 0) || (v1 > 0 && v2 > 0)
	}
	for {
		miny := calc(minx)
		maxy := calc(maxx)
		if ss(miny, maxy) {
			panic(fmt.Errorf("%s: 函数值在'%v'和'%v'两处同号", c.dump(*param), minx, maxx))
		}
		midx := (minx + maxx) / 2
		midy := calc(midx)
		if ss(miny, midy) {
			minx = midx
		} else {
			maxx = midx
		}
	}
}

func (c *Calculation) root(x string) {
	var data [][]float64
	for _, p := range c.Args.Params() {
		c.bisect(x, &p)
		var row []float64
		for _, c := range c.Data {
			v := p[c].(float64)
			row = append(row, v)
		}
		data = append(data, row)
	}
	for _, fn := range c.Save {
		fn = c.fullPath(fn, nil)
		switch filepath.Ext(fn) {
		case ".csv":
			c.saveCSV(fn, data)
		case ".png":
			c.savePNG(fn, data)
		}
	}
}

func (r *Recipe) Solve() (err error) {
	if r.Root == nil {
		return
	}
	defer func() { err = allege(recover()) }()
	if r.Conf.Epsilon <= 0 || r.Conf.Epsilon > 1 {
		panic(errors.New("对分法解方程的允许误差'epsilon'必须是0~1之间的正数"))
	}
	r.Root.err = r.Conf.Epsilon
	r.Root.root(r.Func.IVar)
	return
}
