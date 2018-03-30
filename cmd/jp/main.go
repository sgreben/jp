package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/sgreben/jp/pkg/draw"
	"github.com/sgreben/jp/pkg/terminal"

	"github.com/sgreben/jp/pkg/jsonpath"
)

type configuration struct {
	Box        draw.Box
	X          string
	Y          string
	XY         string
	PlotType   enumVar
	CanvasType enumVar
	InputType  enumVar
	HistBins   uint
}

const (
	plotTypeLine    = "line"
	plotTypeBar     = "bar"
	plotTypeScatter = "scatter"
	plotTypeHist    = "hist"
	plotTypeHist2D  = "hist2d"
)

const (
	canvasTypeFull    = "full"
	canvasTypeQuarter = "quarter"
	canvasTypeBraille = "braille"
	canvasTypeAuto    = "auto"
)

const (
	inputTypeCSV  = "csv"
	inputTypeJSON = "json"
)

var config = configuration{
	PlotType: enumVar{
		Value: plotTypeLine,
		Choices: []string{
			plotTypeLine,
			plotTypeBar,
			plotTypeScatter,
			plotTypeHist,
			plotTypeHist2D,
		},
	},
	CanvasType: enumVar{
		Value: canvasTypeAuto,
		Choices: []string{
			canvasTypeFull,
			canvasTypeQuarter,
			canvasTypeBraille,
			canvasTypeAuto,
		},
	},
	InputType: enumVar{
		Value: inputTypeJSON,
		Choices: []string{
			inputTypeJSON,
			inputTypeCSV,
		},
	},
}

var (
	xPattern  *jsonpath.JSONPath
	yPattern  *jsonpath.JSONPath
	xyPattern *jsonpath.JSONPath
)

func init() {
	flag.Var(&config.PlotType, "type", fmt.Sprintf("Plot type. One of %v", config.PlotType.Choices))
	flag.Var(&config.CanvasType, "canvas", fmt.Sprintf("Canvas type. One of %v", config.CanvasType.Choices))
	flag.Var(&config.InputType, "input", fmt.Sprintf("Input type. One of %v", config.InputType.Choices))
	flag.StringVar(&config.X, "x", "", "x values (JSONPath expression)")
	flag.StringVar(&config.Y, "y", "", "y values (JSONPath expression)")
	flag.StringVar(&config.XY, "xy", "", "x,y value pairs (JSONPath expression). Overrides -x and -y if given.")
	flag.IntVar(&config.Box.Width, "width", 0, "Plot width (default 0 (auto))")
	flag.IntVar(&config.Box.Height, "height", 0, "Plot height (default 0 (auto))")
	flag.UintVar(&config.HistBins, "bins", 0, "Number of histogram bins (default 0 (auto))")
	flag.Parse()
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error
	xPattern = jsonpath.New("x")
	xPattern.AllowMissingKeys(true)
	err = xPattern.Parse(fmt.Sprintf("{%s}", config.X))
	if err != nil {
		log.Fatal(err)
	}
	yPattern = jsonpath.New("y")
	yPattern.AllowMissingKeys(true)
	err = yPattern.Parse(fmt.Sprintf("{%s}", config.Y))
	if err != nil {
		log.Fatal(err)
	}
	if config.XY != "" || (config.X == "" && config.Y == "") {
		xyPattern = jsonpath.New("xy")
		xyPattern.AllowMissingKeys(true)
		err = xyPattern.Parse(fmt.Sprintf("{%s}", config.XY))
		if err != nil {
			log.Fatal(err)
		}
	}
	if config.Box.Width == 0 {
		config.Box.Width = terminal.Width()
	}
	if config.Box.Height == 0 {
		config.Box.Height = terminal.Height() - 1
	}
	if config.CanvasType.Value == canvasTypeAuto {
		config.CanvasType.Value = autoCanvas[config.PlotType.Value]
	}
}

func match(in interface{}, p *jsonpath.JSONPath) [][]reflect.Value {
	defer func() {
		if r := recover(); r != nil {
			log.Println("error evaluating JSONPath", p.String+":", r)
		}
	}()
	out, err := p.FindResults(in)
	if err != nil {
		log.Println(err)
	}
	return out
}

func main() {
	var in interface{}
	switch config.InputType.Value {
	case inputTypeJSON:
		dec := json.NewDecoder(os.Stdin)
		err := dec.Decode(&in)
		if err != nil {
			log.Println(inputTypeJSON, "input:", err)
		}
	case inputTypeCSV:
		r := csv.NewReader(os.Stdin)
		rows, err := r.ReadAll()
		if err != nil {
			log.Println(inputTypeCSV, "input:", err)
		}
		in = parseRows(rows)
	}
	var x, y []reflect.Value
	if xyPattern != nil {
		x, y = split(match(in, xyPattern))
	} else {
		x = flatten(match(in, xPattern))
		y = flatten(match(in, yPattern))
	}
	buffer := draw.NewBuffer(config.Box)
	var p draw.Pixels
	switch config.CanvasType.Value {
	case canvasTypeBraille:
		p = &draw.Braille{Buffer: buffer}
	case canvasTypeQuarter:
		p = &draw.Quarter{Buffer: buffer}
	case canvasTypeFull:
		p = &draw.Full{Buffer: buffer}
	}
	p.Clear()
	c := draw.Canvas{Pixels: p}
	switch config.PlotType.Value {
	case plotTypeLine:
		fmt.Println(linePlot(x, y, c))
	case plotTypeScatter:
		fmt.Println(scatterPlot(x, y, c))
	case plotTypeBar:
		fmt.Println(barPlot(x, y, c))
	case plotTypeHist:
		fmt.Println(histogram(x, c, config.HistBins))
	case plotTypeHist2D:
		fmt.Println(hist2D(x, y, c, config.HistBins))
	}
}
