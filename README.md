# catmullrom

Package catmullrom provides an implementation of the centripetal
Catmull-Rom spline. It calculates points of a 2D spline curve given a series
of control points.

![Catmull-Rom spline](demo/points.png?raw=true "Catmull-Rom spline")

Add it to a module as a dependency via:

    go get github.com/fzipp/catmullrom

## Example usage

    package main

    import (
    	"fmt"
    
    	"github.com/fzipp/catmullrom"
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
    
    	curve := catmullrom.SplineChain(controlPoints, 100, 0.5)
    
    	for _, point := range curve {
    		fmt.Printf("%v\t%v\n", point.X, point.Y)
    	}
    }

