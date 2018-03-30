package main

import (
	"log"
	"reflect"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
	"github.com/sgreben/jp/pkg/plot"
)

func heatmapData(xv []reflect.Value, yv []reflect.Value, nbins uint) (heatmap *data.Heatmap) {
	var x, y []float64
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
	if len(x) != len(y) {
		log.Fatal(len(x), " = len(x) != len(y) = ", len(y))
	}
	points := make([][2]float64, len(x))
	for i := 0; i < len(x); i++ {
		points[i] = [2]float64{x[i], y[i]}
	}
	if len(x) == 0 {
		log.Fatal("no valid x values given")
	}
	bins := data.NewBins2D(points)
	bins.X.Number = int(nbins)
	bins.Y.Number = int(nbins)
	if nbins == 0 {
		bins.X.Number = data.BinsSturges(len(points))
		bins.Y.Number = data.BinsSturges(len(points))
	}
	heatmap = data.NewHeatmap(data.Histogram2D(points, bins))
	return
}

func heatmap(xv, yv []reflect.Value, c draw.Canvas, nbins uint) string {
	heatmap := heatmapData(xv, yv, nbins)
	chart := plot.NewHeatMap(c.GetBuffer())
	return chart.Draw(heatmap)
}
