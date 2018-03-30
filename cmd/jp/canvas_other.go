// +build !windows

package main

var autoCanvas = map[string]string{
	plotTypeBar:     canvasTypeQuarter,
	plotTypeLine:    canvasTypeQuarter,
	plotTypeScatter: canvasTypeBraille,
	plotTypeHist:    canvasTypeQuarter,
	plotTypeHeatmap: canvasTypeFull,
}
