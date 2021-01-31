# catmullrom

[![PkgGoDev](https://pkg.go.dev/badge/github.com/fzipp/catmullrom)](https://pkg.go.dev/github.com/fzipp/catmullrom)
[![Go Report Card](https://goreportcard.com/badge/github.com/fzipp/catmullrom)](https://goreportcard.com/report/github.com/fzipp/catmullrom)

Package catmullrom implements the centripetal Catmull-Rom spline and
calculates points of a 2D spline curve given a series of control points.

Add it to a module as a dependency via:

```
go get github.com/fzipp/catmullrom
```

## Example usage

### Single spline

A single spline requires four control points (p0, p1, p2, p3). The resulting
interpolated points of the spline curve are between the second and the third
control point (p1 and p2).

```go
package main

import (
	"fmt"

	"github.com/fzipp/catmullrom"
)

func main() {

	p := []catmullrom.Point{
		{X: 0, Y: 2.5},
		{X: 2, Y: 4},
		{X: 3, Y: 2},
		{X: 4, Y: 1.5},
	}

	curve := catmullrom.Spline(p[0], p[1], p[2], p[3], 100, 0.5)

	for _, point := range curve {
		fmt.Printf("%v\t%v\n", point.X, point.Y)
	}
}
```

Output:

```
2	4
2.014365414796725	3.9961683757388826
2.0285086520907125	3.991519010789197
2.0424338917428546	3.986066863428536
[...]
2.9631689321190793	2.0405052334680907
2.9752654723162837	2.026380248925892
2.9875411016562947	2.012873518343997
3	2
```

### Spline chain

A spline chain goes through an arbitrary number of control points, including
the first and the last control point.

```go
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
```

Output:

```
0	2.5
0.020250950290489646	2.5153391496514774
0.040597763611450566	2.531045909751864
0.06103744424317995	2.5471087924746265
[...]
7.030349491470336	5.5157336988112
7.0202228847077075	5.5103624536889235
7.01010628000193	5.505116539771162
7	5.5
```

Visualized curve:

![Catmull-Rom spline chain](doc/spline_chain.png?raw=true "Catmull-Rom spline chain")

## License

This project is free and open source software licensed under the
[BSD 3-Clause License](LICENSE).
