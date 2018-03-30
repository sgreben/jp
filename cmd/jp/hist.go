package main

import (
	"log"
	"reflect"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
	"github.com/sgreben/jp/pkg/plot"
)

func histogramData(xv []reflect.Value, nbins uint) (groups []string, counts []float64) {
	var x []float64
	for i := range xv {
		if xv[i].IsValid() && xv[i].CanInterface() {
			xvi, ok := xv[i].Interface().(float64)
			if ok {
				x = append(x, xvi)
			}
		}
	}
	if len(x) == 0 {
		log.Fatal("no valid x values given")
	}
	bins := data.NewBins(x)
	bins.Number = int(nbins)
	if nbins == 0 {
		bins.ChooseSturges()
	}
	hist := data.Histogram(x, bins)
	groups = make([]string, len(hist))
	counts = make([]float64, len(hist))
	for i, b := range hist {
		groups[i] = b.String()
		counts[i] = float64(b.Count)
	}
	return
}

func histogram(xv []reflect.Value, c draw.Canvas, nbins uint) string {
	groups, counts := histogramData(xv, nbins)
	chart := plot.NewBarChart(c)
	chart.BarPaddingX = 0
	data := new(data.Table)
	for _, g := range groups {
		data.AddColumn(g)
	}
	data.AddRow(counts...)
	return chart.Draw(data)
}
