// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/fzipp/catmullrom"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {

	controlPoints := []catmullrom.Point{
		{X: 0, Y: 2.5},
		{X: 2, Y: 4},
		{X: 3, Y: 2},
		{X: 4, Y: 1.5},
		{X: 5, Y: 6},
		{X: 6, Y: 5},
		{X: 7, Y: 3},
		{X: 9, Y: 1},
		{X: 10, Y: 2.5},
		{X: 11, Y: 7},
		{X: 9, Y: 5},
		{X: 8, Y: 6},
		{X: 7, Y: 5.5},
	}

	curve := catmullrom.ChainComplete(controlPoints, 100, 0.5)

	for _, point := range curve {
		fmt.Printf("%v\t%v\n", point.X, point.Y)
	}

	savePlot(curve, controlPoints, "points.png")
}

type xyPoints []catmullrom.Point

func (points xyPoints) Len() int {
	return len(points)
}

func (points xyPoints) XY(i int) (float64, float64) {
	return points[i].X, points[i].Y
}

func savePlot(splinePoints, controlPoints xyPoints, file string) {
	plot.DefaultFont = "Helvetica"
	p, err := plot.New()
	if err != nil {
		log.Fatalln("Could not create plot.", err)
	}

	p.Title.Text = "Catmull-Rom Spline"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s1, err := plotter.NewLine(splinePoints)
	if err != nil {
		log.Fatalln("Could not create scatter plot for spline points.", err)
	}
	s1.Color = plotutil.Color(1)

	s2, err := plotter.NewScatter(controlPoints)
	if err != nil {
		log.Fatalln("Could not create scatter plot for control points.", err)
	}
	s2.Color = plotutil.Color(0)
	s2.Shape = draw.CircleGlyph{}

	p.Add(s1, s2)

	if err := p.Save(16*vg.Centimeter, 12*vg.Centimeter, file); err != nil {
		log.Fatalln("Could not save plot to image file.", err)
	}
}
