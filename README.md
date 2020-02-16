# catmullrom

Package catmullrom implements the centripetal Catmull-Rom spline and
calculates points of a 2D spline curve given a series of control points.

Add it to a module as a dependency via:

```
go get github.com/fzipp/catmullrom
```

## Documentation

Package documentation can be [found on pkg.go.dev](https://pkg.go.dev/github.com/fzipp/catmullrom?tab=doc).

## Example usage

### Spline chain

```
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

![Catmull-Rom spline chain](doc/spline_chain.png?raw=true "Catmull-Rom spline chain")


### Single spline

```
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

## License

This project is free and open source software licensed under the
[BSD 3-Clause License](LICENSE).
