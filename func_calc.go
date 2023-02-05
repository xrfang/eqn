package main

import (
	"fmt"
	"path/filepath"
)

func (c *Calculation) calc(x string) {
	params := []map[string]any{{}}
	if len(c.Args) > 0 {
		params = c.Args.Params()
	}
	for _, p := range params {
		var data [][]float64
		for _, iv := range c.IVar {
			p[x] = iv
			c.call(&p)
			var row []float64
			for _, col := range c.Data {
				v, ok := p[col].(float64)
				if !ok {
					panic(fmt.Errorf("变量'%s'在结果中不存在", col))
				}
				row = append(row, v)
			}
			data = append(data, row)
		}
		for _, fn := range c.Save {
			fn = c.fullPath(fn, p)
			switch filepath.Ext(fn) {
			case ".csv":
				c.saveCSV(fn, data)
			case ".png":
				c.savePNG(fn, data)
			}
		}
	}
}

func (r *Recipe) Calculate() (err error) {
	if r.Calc == nil {
		return
	}
	defer func() { err = allege(recover()) }()
	r.Calc.calc(r.Func.IVar)
	return
}
