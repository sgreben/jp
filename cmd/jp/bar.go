package main

import (
	"fmt"
	"reflect"

	"github.com/sgreben/jp/pkg/jp"
	"github.com/sgreben/jp/pkg/jp/primitives"
)

func barPlotData(xvv, yvv [][]reflect.Value) (x []string, y []float64) {
	for _, xv := range xvv {
		for i := range xv {
			x = append(x, fmt.Sprint(xv[i].Interface()))
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

func barPlot(xvv, yvv [][]reflect.Value, box primitives.Box) string {
	groups, y := barPlotData(xvv, yvv)
	chart := jp.NewBarChart(box.Width, box.Height)
	data := new(jp.DataTable)
	if len(groups) != len(y) {
		for i := range y {
			data.AddColumn(fmt.Sprint(i))
		}
	} else {
		for _, g := range groups {
			data.AddColumn(g)
		}
	}
	data.AddRow(y...)
	return chart.Draw(data)
}
