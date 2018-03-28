package main

import (
	"fmt"
	"reflect"

	"github.com/sgreben/jp/pkg/draw"
	"github.com/sgreben/jp/pkg/plot"
)

func barPlotData(xvv, yvv [][]reflect.Value) (x []string, y []float64) {
	for _, xv := range xvv {
		for i := range xv {
			if xv[i].IsValid() && xv[i].CanInterface() {
				x = append(x, fmt.Sprint(xv[i].Interface()))
			}
		}
	}
	for _, yv := range yvv {
		for i := range yv {
			if yv[i].IsValid() && yv[i].CanInterface() {
				yvi, ok := yv[i].Interface().(float64)
				if ok {
					y = append(y, yvi)
				}
			}
		}
	}
	return
}

func barPlot(xvv, yvv [][]reflect.Value, c draw.Canvas) string {
	groups, y := barPlotData(xvv, yvv)
	chart := plot.NewBarChart(c)
	data := new(plot.DataTable)
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
