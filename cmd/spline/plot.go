package main

import (
	"fmt"

	"github.com/fzipp/catmullrom"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type xyPoints []catmullrom.Point

func (points xyPoints) Len() int {
	return len(points)
}

func (points xyPoints) XY(i int) (float64, float64) {
	return points[i].X, points[i].Y
}

func savePlot(splinePoints, controlPoints xyPoints, file string) error {
	plot.DefaultFont = "Helvetica"
	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %w", err)
	}

	p.Title.Text = "Catmull-Rom Spline"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s1, err := plotter.NewLine(splinePoints)
	if err != nil {
		return fmt.Errorf("could not create scatter plot for spline points: %w", err)
	}
	s1.Color = plotutil.Color(1)

	s2, err := plotter.NewScatter(controlPoints)
	if err != nil {
		return fmt.Errorf("could not create scatter plot for control points: %w", err)
	}
	s2.Color = plotutil.Color(0)
	s2.Shape = draw.CircleGlyph{}

	p.Add(s1, s2)

	if err := p.Save(16*vg.Centimeter, 12*vg.Centimeter, file); err != nil {
		return fmt.Errorf("could not save plot to image file: %w", err)
	}
	return nil
}
