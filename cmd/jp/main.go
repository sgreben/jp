package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"

	"github.com/mattn/go-runewidth"
	"github.com/sgreben/jp/pkg/jp/primitives"
	"github.com/sgreben/jp/pkg/terminal"

	"github.com/sgreben/jp/pkg/jsonpath"
)

type configuration struct {
	Box      primitives.Box
	X        string
	Y        string
	XY       string
	PlotType enumVar
}

const plotTypeLine = "line"
const plotTypeBar = "bar"
const plotTypeScatter = "scatter"
const plotTypeHist = "hist"

var config = configuration{
	PlotType: enumVar{
		Value: plotTypeLine,
		Choices: []string{
			plotTypeLine,
			plotTypeBar,
		},
	},
}

var xPattern *jsonpath.JSONPath
var yPattern *jsonpath.JSONPath
var xyPattern *jsonpath.JSONPath

func init() {
	flag.Var(&config.PlotType, "type", fmt.Sprintf("Plot type. One of %v", config.PlotType.Choices))
	flag.StringVar(&config.X, "x", "", "x values (JSONPath expression)")
	flag.StringVar(&config.Y, "y", "", "y values (JSONPath expression)")
	flag.StringVar(&config.XY, "xy", "", "x,y value pairs (JSONPath expression). Overrides -x and -y if given.")
	flag.IntVar(&config.Box.Width, "width", 0, "Plot width (default 0 (auto))")
	flag.IntVar(&config.Box.Height, "height", 0, "Plot height (default 0 (auto))")
	flag.Parse()
	log.SetOutput(os.Stderr)

	var err error
	xPattern = jsonpath.New("x")
	err = xPattern.Parse(fmt.Sprintf("{%s}", config.X))
	if err != nil {
		log.Fatal(err)
	}
	yPattern = jsonpath.New("y")
	err = yPattern.Parse(fmt.Sprintf("{%s}", config.Y))
	if err != nil {
		log.Fatal(err)
	}
	if config.XY != "" {
		xyPattern = jsonpath.New("xy")
		err = xyPattern.Parse(fmt.Sprintf("{%s}", config.XY))
		if err != nil {
			log.Fatal(err)
		}
	}
	if config.Box.Width == 0 {
		config.Box.Width = terminal.Width()
		if runtime.GOOS == "windows" {
			config.Box.Width--
		}
	}
	if config.Box.Height == 0 {
		config.Box.Height = terminal.Height() - 1
	}
	if runewidth.StringWidth(primitives.HorizontalLine) == 2 {
		primitives.HorizontalLine = "-"
		primitives.VerticalLine = "|"
		primitives.CornerBottomLeft = "+"
		primitives.PointSymbolDefault = "*"
		primitives.Cross = "X"
	}
}

func match(in interface{}, p *jsonpath.JSONPath) [][]reflect.Value {
	out, err := p.FindResults(in)
	if err != nil {
		log.Println(err)
	}
	return out
}

func main() {
	var in interface{}
	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(&in)
	if err != nil {
		log.Println(err)
	}
	fmt.Println()
	var x, y [][]reflect.Value
	if xyPattern != nil {
		x, y = split(match(in, xyPattern))
	} else {
		x = match(in, xPattern)
		y = match(in, yPattern)
	}
	switch config.PlotType.Value {
	case plotTypeLine:
		fmt.Print(linePlot(x, y, config.Box))
	case plotTypeBar:
		fmt.Print(barPlot(x, y, config.Box))
	}

}
