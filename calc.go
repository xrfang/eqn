package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Calculation struct {
	IVar Range    `yaml:"ivar"`
	Args Ranges   `yaml:"args"`
	Data []string `yaml:"data"`
	Save []string `yaml:"save"`
	expr []Expression
	err  float64
	min  float64
	max  float64
	nfmt string //浮点数格式
	wdir string //工作目录
}

func (c *Calculation) Validate(r *Recipe) {
	assert(len(c.IVar) > 0, "自变量范围未指定")
	assert(len(c.Data) > 0, "输出数据列未指定")
	assert(len(c.Save) > 0, "输出文件未指定")
	for i, s := range c.Save {
		s = strings.ToLower(s)
		if !strings.HasSuffix(s, ".csv") && !strings.HasSuffix(s, ".png") {
			panic(fmt.Errorf("输出文件名'%s'不正确，必须是'.csv'或'.png'", s))
		}
		c.Save[i] = s
	}
	c.min = c.IVar[0]
	c.max = c.IVar[len(c.IVar)-1]
	c.expr = r.Func.expr
	if r.Conf.Precision > 0 {
		c.nfmt = fmt.Sprintf("%%0.%df", r.Conf.Precision)
	} else {
		c.nfmt = "%e"
	}
	c.wdir = r.path
}

func (c *Calculation) fullPath(fn string, param map[string]any) string {
	if len(param) > 0 {
		for a := range c.Args {
			key := fmt.Sprintf("${%s}", a)
			val := fmt.Sprintf(c.nfmt, param[a])
			fn = strings.ReplaceAll(fn, key, val)
		}
	}
	return filepath.Join(c.wdir, fn)
}

func (c *Calculation) saveCSV(fn string, data [][]float64) {
	f, err := os.Create(fn)
	assert(err)
	defer func() { assert(f.Close()) }()
	for _, row := range data {
		var cols []string
		for _, r := range row {
			cols = append(cols, fmt.Sprintf(c.nfmt, r))
		}
		fmt.Fprintln(f, strings.Join(cols, ","))
	}
}

func (c *Calculation) savePNG(fn string, data [][]float64) {
	var xys plotter.XYs
	p := plot.New()
	p.X.Max = -1e300
	p.X.Min = 1e300
	p.Y.Max = -1e300
	p.Y.Min = 1e300
	for _, r := range data {
		xy := plotter.XY{X: r[0], Y: r[len(r)-1]}
		if err := plotter.CheckFloats(xy.X, xy.Y); err != nil {
			panic(fmt.Errorf("savePNG: x=%v; y=%v; err=%v", xy.X, xy.Y, err))
		}
		if p.X.Min > xy.X {
			p.X.Min = xy.X
		}
		if p.X.Max < xy.X {
			p.X.Max = xy.X
		}
		if p.Y.Min > xy.Y {
			p.Y.Min = xy.Y
		}
		if p.Y.Max < xy.Y {
			p.Y.Max = xy.Y
		}
		xys = append(xys, xy)
	}
	p.X.Label.Text = c.Data[0]
	p.Y.Label.Text = c.Data[len(c.Data)-1]
	g := plotter.NewGrid()
	p.Add(g)
	line, err := plotter.NewLine(xys)
	assert(err)
	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)
	p.Add(line)
	p.Save(imgW, imgH, fn)
}

func (c *Calculation) call(p *map[string]any) (float64, string) {
	var rk string
	var rv float64
	for _, e := range c.expr {
		res, err := e.v.Evaluate(*p)
		assert(err)
		rk = e.n
		rv = res.(float64)
		(*p)[rk] = rv
	}
	return rv, rk
}
