package main

import (
	"github.com/fstanis/screenresolution"
	"gonum.org/v1/plot/font"
)

var imgW, imgH font.Length

func init() {
	imgW = 1280
	imgH = 800
	scale := 0.5
	res := screenresolution.GetPrimary()
	if res != nil {
		imgW = font.Length(float64(res.Width) * scale)
		imgH = font.Length(float64(res.Height) * scale)
	}
}
