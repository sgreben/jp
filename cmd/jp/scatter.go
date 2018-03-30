package main

import (
	"log"
	"reflect"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
	"github.com/sgreben/jp/pkg/plot"
)

func scatterPlotData(xv, yv []reflect.Value) (x, y []float64) {
	for i := range xv {
		if xv[i].IsValid() && xv[i].CanInterface() {
			xvi, ok := xv[i].Interface().(float64)
			if ok {
				x = append(x, xvi)
			}
		}
	}
	for i := range yv {
		if yv[i].IsValid() && yv[i].CanInterface() {
			yvi, ok := yv[i].Interface().(float64)
			if ok {
				y = append(y, yvi)
			}
		}
	}
	return
}

func scatterPlot(xv, yv []reflect.Value, c draw.Canvas) string {
	x, y := scatterPlotData(xv, yv)
	chart := plot.NewScatterChart(c)
	data := new(data.Table)
	data.AddColumn("x")
	data.AddColumn("y")
	n := len(x)
	if len(y) > n {
		n = len(y)
	}
	if len(y) == 0 {
		log.Fatal("no valid y values given")
	}
	// If no valid xs are given, use the indices as x values.
	if len(x) == 0 {
		x = make([]float64, len(y))
		for i := 0; i < len(y); i++ {
			x[i] = float64(i)
		}
	}
	for i := 0; i < n; i++ {
		data.AddRow(x[i%len(x)], y[i%len(y)])
	}
	return chart.Draw(data)
}
