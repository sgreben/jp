package plot

import (
	"bytes"
	"math"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
)

// BarChart is a bar chart
type BarChart struct {
	draw.Canvas
	BarPaddingX int
}

// NewBarChart returns a new bar chart
func NewBarChart(canvas draw.Canvas) *BarChart {
	chart := new(BarChart)
	chart.Canvas = canvas
	chart.BarPaddingX = 1
	return chart
}

// Draw implements Chart
func (c *BarChart) Draw(table *data.Table) string {
	minY := math.Inf(1)
	maxY := math.Inf(-1)
	for _, row := range table.Rows {
		for _, y := range row {
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	paddingX := 2
	paddingY := 3
	chartHeight := c.Size().Height - paddingY*c.RuneSize().Height
	chartWidth := c.Size().Width - 2*paddingX*c.RuneSize().Width

	labelsBelowBars := true
	labelsRight := false
	maxLabelLength := 0
	totalLabelLength := 0
	for _, group := range table.Columns {
		totalLabelLength += len(group)
		if len(group) > maxLabelLength {
			maxLabelLength = len(group)
		}
	}
	if totalLabelLength*c.RuneSize().Width > chartWidth {
		labelsBelowBars = false
		if len(table.Columns)*c.RuneSize().Height <= chartHeight {
			labelsRight = true
			chartWidth -= 3 + maxLabelLength*c.RuneSize().Width
		}
	}

	scaleY := float64(chartHeight) / maxY
	barPaddedWidth := chartWidth / len(table.Columns)
	barWidth := barPaddedWidth - (c.BarPaddingX * c.RuneSize().Width)
	if barPaddedWidth < c.RuneSize().Width {
		barPaddedWidth = c.RuneSize().Width
	}
	if barWidth < c.RuneSize().Width {
		barWidth = c.RuneSize().Width
	}

	scaleY = float64(chartHeight) / maxY

	for i, group := range table.Columns {
		barLeft := paddingX*c.RuneSize().Width + barPaddedWidth*i
		barRight := barLeft + barWidth

		y := table.Rows[0][i]
		barHeight := y * scaleY
		barBottom := (paddingY - 1) * c.RuneSize().Height
		barTop := barBottom + int(barHeight)

		for x := barLeft; x < barRight; x++ {
			for y := barBottom; y < barTop; y++ {
				c.Set(y, x)
			}
		}

		barMiddle := int(math.Floor(float64(barLeft+barRight) / float64(2*c.RuneSize().Width)))

		// Group label
		if labelsBelowBars {
			c.GetBuffer().WriteCenter(0, barMiddle, []rune(group))
		} else {
			c.GetBuffer().WriteCenter(0, barMiddle, []rune(Fi(i)))
		}

		// Count label
		countLabelY := int(math.Ceil(float64(barTop)/float64(c.RuneSize().Height))) * c.RuneSize().Height
		if countLabelY <= barBottom && y > 0 {
			c.GetBuffer().SetRow(barTop/c.RuneSize().Height, barLeft/c.RuneSize().Width, barRight/c.RuneSize().Width, '▁')
			countLabelY = 3 * c.RuneSize().Height
		}
		c.GetBuffer().WriteCenter(countLabelY/c.RuneSize().Height, barMiddle, Ff(y))
	}
	if labelsRight {
		for i, group := range table.Columns {
			c.GetBuffer().WriteRight(c.GetBuffer().Height-i, paddingX+1+chartWidth/c.RuneSize().Width, []rune(Fi(i)))
			c.GetBuffer().WriteRight(c.GetBuffer().Height-i, paddingX+4+chartWidth/c.RuneSize().Width, []rune(group))
		}
	}

	b := bytes.NewBuffer(nil)
	c.GetBuffer().Render(b)
	return b.String()
}
