package main

import (
	"reflect"

	"github.com/sgreben/jp/pkg/jp"
	"github.com/sgreben/jp/pkg/jp/primitives"
)

func linePlotData(xvv, yvv [][]reflect.Value) (x, y []float64) {
	for _, xv := range xvv {
		for i := range xv {
			xvi, ok := xv[i].Interface().(float64)
			if ok {
				x = append(x, xvi)
			}
		}
	}
	for _, yv := range yvv {
		for i := range yv {
			yvi, ok := yv[i].Interface().(float64)
			if ok {
				y = append(y, yvi)
			}
		}
	}
	return
}

func linePlot(xvv, yvv [][]reflect.Value, box primitives.Box) string {
	x, y := linePlotData(xvv, yvv)
	chart := jp.NewLineChart(box.Width, box.Height)
	data := new(jp.DataTable)
	data.AddColumn("x")
	data.AddColumn("y")
	n := len(x)
	if len(y) > n {
		n = len(y)
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
	chart.Symbol = "█"
	return chart.Draw(data)
}
