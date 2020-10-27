package main

import (
	"fmt"
	"log"
	"math"
	"reflect"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
	"github.com/sgreben/jp/pkg/plot"
)

func barPlotData(xv, yv []reflect.Value, digits int) (x []string, y []float64) {
	for i := range xv {
		if xv[i].IsValid() && xv[i].CanInterface() {
			x = append(x, fmt.Sprint(xv[i].Interface()))
		}
	}
	for i := range yv {
		if yv[i].IsValid() && yv[i].CanInterface() {
			yvi, ok := yv[i].Interface().(float64)
			if ok {
				y = append(y, math.Round(yvi*math.Pow10(digits))/math.Pow10(digits))
			}
		}
	}
	return
}

func barPlot(xv, yv []reflect.Value, c draw.Canvas, digits int) string {
	groups, y := barPlotData(xv, yv, digits)
	chart := plot.NewBarChart(c)
	data := new(data.Table)
	if len(y) == 0 {
		log.Fatal("no valid y values given")
	}
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
