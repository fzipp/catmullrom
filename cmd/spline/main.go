// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This little tool visualizes a Catmull-Rom spline chain for a series
// of control points by creating a plot PNG image of the resulting curve.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/fzipp/catmullrom"
)

func usage() {
	fail(`usage: spline [-pps n] [-alpha a] [-plot png_output_file] control_points_file`)
}

func main() {
	pointsPerSegment := flag.Int("pps", 100, "")
	alpha := flag.Float64("alpha", 0.5, "")
	plotFilePath := flag.String("plot", "", "")

	flag.Usage = usage
	flag.Parse()

	var err error
	inputFile := os.Stdin
	if flag.NArg() > 0 {
		inputFile, err = os.Open(flag.Arg(0))
		check(err)
		defer inputFile.Close()
	}

	controlPoints, err := parsePoints(inputFile)
	check(err)

	curve := catmullrom.SplineChain(controlPoints, *pointsPerSegment, *alpha)

	for _, point := range curve {
		fmt.Printf("%v\t%v\n", point.X, point.Y)
	}

	if *plotFilePath != "" {
		err := savePlot(curve, controlPoints, *plotFilePath)
		check(err)
	}
}

func parsePoints(r io.Reader) ([]catmullrom.Point, error) {
	var points []catmullrom.Point
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}
		x, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid X coordinate: '%s'", fields[0])
		}
		y, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid Y coordinate: '%s'", fields[1])
		}
		points = append(points, catmullrom.Point{X: x, Y: y})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read from input: %w", err)
	}
	return points, nil
}

func check(err error) {
	if err != nil {
		fail(err)
	}
}

func fail(message interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
