package jp

import (
	"math"
	"strings"

	"github.com/sgreben/jp/pkg/jp/primitives"
)

// BarChart is a bar chart
type BarChart struct {
	Buffer []string

	data *DataTable

	Width       int
	Height      int
	Symbol      string
	BarPaddingX int

	chartHeight int
	chartWidth  int

	paddingX int
	paddingY int
}

// NewBarChart returns a new bar chart
func NewBarChart(width, height int) *BarChart {
	chart := new(BarChart)
	chart.Width = width
	chart.Height = height
	chart.Buffer = primitives.Buffer(width * height)
	chart.Symbol = primitives.FullBlock4
	chart.paddingY = 3
	chart.BarPaddingX = 1
	return chart
}

func (c *BarChart) writeText(text string, x, y int) {
	coord := y*c.Width + x

	for idx, char := range strings.Split(text, "") {
		if coord+idx >= 0 && coord+idx < len(c.Buffer) {
			c.Buffer[coord+idx] = char
		}
	}
}

// Draw implements Chart
func (c *BarChart) Draw(data *DataTable) (out string) {

	c.data = data

	minY := math.Inf(1)
	maxY := math.Inf(-1)
	for _, row := range data.Rows {
		for _, y := range row {
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	c.paddingX = 1
	c.chartHeight = c.Height - c.paddingY
	c.chartWidth = c.Width - c.paddingX - 1
	scaleY := float64(c.chartHeight) / maxY
	barPaddedWidth := c.chartWidth / len(data.Columns)
	barWidth := barPaddedWidth - c.BarPaddingX
	if barPaddedWidth < 1 {
		barPaddedWidth = 1
	}
	if barWidth < 1 {
		barWidth = 1
	}

	scaleY = float64(c.chartHeight) / maxY

	for i, group := range data.Columns {
		barLeft := c.paddingX + barPaddedWidth*i
		barRight := barLeft + barWidth
		y := data.Rows[0][i]
		barHeight := y * scaleY
		barTop := c.paddingY - 1 + int(barHeight)
		for x := barLeft; x < barRight; x++ {
			for y := c.paddingY - 1; y < barTop; y++ {
				c.set(x, y, c.Symbol)
			}
			if barHeight < 1 && y > 0 {
				for x := barLeft; x < barRight; x++ {
					c.set(x, barTop, "▁")
				}
			}
		}

		// Group label
		barMiddle := (barLeft + barRight) / 2
		c.writeText(group, barMiddle-(len(group)/2), 0)

		// Count label
		countLabelY := barTop
		if barHeight < 1 {
			countLabelY = barTop + 1
		}
		c.writeText(primitives.Ff(y), barMiddle-(len(primitives.Ff(y))/2), countLabelY)
	}

	for row := c.Height - 1; row >= 0; row-- {
		out += strings.Join(c.Buffer[row*c.Width:(row+1)*c.Width], "") + "\n"
	}

	return
}

func (c *BarChart) set(x, y int, s string) {
	coord := y*c.Width + x
	if coord > 0 && coord < len(c.Buffer) {
		c.Buffer[coord] = s
	}
}
