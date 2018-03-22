package jp

import (
	"math"
	"strings"

	"github.com/sgreben/jp/pkg/jp/primitives"
)

// Adapted from https://github.com/buger/goterm under the MIT License.

/*
MIT License

Copyright (c) 2016 Leonid Bugaev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

// LineChart is a line chart
type LineChart struct {
	Buffer      []string
	Width       int
	Height      int
	Symbol      string
	chartHeight int
	chartWidth  int
	paddingX    int
	paddingY    int

	data *DataTable
}

// NewLineChart returns a new line chart
func NewLineChart(width, height int) *LineChart {
	chart := new(LineChart)
	chart.Width = width
	chart.Height = height
	chart.Buffer = primitives.Buffer(width * height)
	chart.Symbol = primitives.PointSymbolDefault
	chart.paddingY = 2
	return chart
}

func (c *LineChart) drawAxes(maxX, minX, maxY, minY float64, index int) {
	c.drawLine(c.paddingX-1, 1, c.Width-1, 1, primitives.HorizontalLine)
	c.drawLine(c.paddingX-1, 1, c.paddingX-1, c.Height-1, primitives.VerticalLine)
	c.set(c.paddingX-1, c.paddingY-1, primitives.CornerBottomLeft)

	left := 0 // c.Width - c.paddingX + 1
	c.writeText(primitives.Ff(minY), left, 1)

	c.writeText(primitives.Ff(maxY), left, c.Height-1)

	c.writeText(primitives.Ff(minX), c.paddingX, 0)

	xCol := c.data.Columns[0]
	c.writeText(c.data.Columns[0], c.Width/2-len(xCol)/2, 1)

	if len(c.data.Columns) < 3 {
		col := c.data.Columns[index]

		for idx, char := range strings.Split(col, "") {
			startFrom := c.Height/2 + len(col)/2 - idx

			c.writeText(char, c.paddingX-1, startFrom)
		}
	}

	c.writeText(primitives.Ff(maxX), c.Width-len(primitives.Ff(maxX)), 0)

}

func (c *LineChart) writeText(text string, x, y int) {
	coord := y*c.Width + x

	for idx, char := range strings.Split(text, "") {
		c.Buffer[coord+idx] = char
	}
}

// Draw implements Chart
func (c *LineChart) Draw(data *DataTable) (out string) {
	var scaleY, scaleX float64

	c.data = data

	charts := len(data.Columns) - 1

	prevPoint := [2]int{-1, -1}

	maxX, minX, maxY, minY := getBoundaryValues(data, -1)

	c.paddingX = int(math.Max(float64(len(primitives.Ff(minY))), float64(len(primitives.Ff(maxY))))) + 1

	c.chartHeight = c.Height - c.paddingY
	c.chartWidth = c.Width - c.paddingX - 1

	scaleX = float64(c.chartWidth) / (maxX - minX)

	scaleY = float64(c.chartHeight) / (maxY - minY)

	for i := 1; i < charts+1; i++ {
		symbol := c.Symbol

		chartData := getChartData(data, i)

		for _, point := range chartData {
			x := int((point[0]-minX)*scaleX) + c.paddingX
			y := int((point[1])*scaleY) + c.paddingY
			y = int((point[1]-minY)*scaleY) + c.paddingY

			if prevPoint[0] == -1 {
				prevPoint[0] = x
				prevPoint[1] = y
			}

			if prevPoint[0] <= x {
				c.drawLine(prevPoint[0], prevPoint[1], x, y, symbol)
			}

			prevPoint[0] = x
			prevPoint[1] = y
		}

		c.drawAxes(maxX, minX, maxY, minY, i)
	}

	for row := c.Height - 1; row >= 0; row-- {
		out += strings.Join(c.Buffer[row*c.Width:(row+1)*c.Width], "") + "\n"
	}

	return
}

func (c *LineChart) set(x, y int, s string) {
	coord := y*c.Width + x
	if coord > 0 && coord < len(c.Buffer) {
		c.Buffer[coord] = s
	}
}

func (c *LineChart) drawLine(x0, y0, x1, y1 int, symbol string) {
	primitives.DrawLine(x0, y0, x1, y1, func(x, y int) { c.set(x, y, symbol) })
}

func getBoundaryValues(data *DataTable, index int) (maxX, minX, maxY, minY float64) {
	maxX = math.Inf(-1)
	minX = math.Inf(1)
	maxY = math.Inf(-1)
	minY = math.Inf(1)

	for _, r := range data.Rows {
		maxX = math.Max(maxX, r[0])
		minX = math.Min(minX, r[0])

		for idx, c := range r {
			if idx > 0 {
				if index == -1 || index == idx {
					maxY = math.Max(maxY, c)
					minY = math.Min(minY, c)
				}
			}
		}
	}

	if maxY > 0 {
		maxY = maxY * 1.1
	} else {
		maxY = maxY * 0.9
	}

	if minY > 0 {
		minY = minY * 0.9
	} else {
		minY = minY * 1.1
	}

	return
}

func getChartData(data *DataTable, index int) (out [][]float64) {
	for _, r := range data.Rows {
		out = append(out, []float64{r[0], r[index]})
	}

	return
}
