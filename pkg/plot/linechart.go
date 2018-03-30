package plot

import (
	"bytes"
	"math"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
)

// LineChart is a line chart
type LineChart struct{ draw.Canvas }

// NewLineChart returns a new line chart
func NewLineChart(canvas draw.Canvas) *LineChart { return &LineChart{canvas} }

func (c *LineChart) drawAxes(paddingX, paddingY int, minX, maxX, minY, maxY float64) {
	buffer := c.GetBuffer()
	// X axis
	buffer.SetRow(1, paddingX, buffer.Width, draw.HorizontalLine)
	// Y axis
	buffer.SetColumn(1, buffer.Height, paddingX, draw.VerticalLine)
	// Corner
	buffer.Set(1, paddingX, draw.CornerBottomLeft)
	// Labels
	buffer.WriteRight(1, 1, Ff(minY))
	buffer.WriteLeft(buffer.Height-1, paddingX, Ff(maxY))
	buffer.WriteRight(0, paddingX, Ff(minX))
	buffer.WriteLeft(0, buffer.Width, Ff(maxX))
}

// Draw implements Chart
func (c *LineChart) Draw(table *data.Table) string {
	var scaleY, scaleX float64
	var prevX, prevY int

	minX, maxX, minY, maxY := minMax(table)
	minLabelWidth := len(Ff(minY))
	maxLabelWidth := len(Ff(maxY))

	paddingX := minLabelWidth + 1
	paddingY := 2
	if minLabelWidth < maxLabelWidth {
		paddingX = maxLabelWidth + 1
	}
	chartWidth := c.Size().Width - (paddingX+1)*c.RuneSize().Width
	chartHeight := c.Size().Height - paddingY*c.RuneSize().Height
	scaleX = float64(chartWidth) / (maxX - minX)
	scaleY = float64(chartHeight) / (maxY - minY)

	first := true
	for _, point := range table.Rows {
		if len(point) < 2 {
			continue
		}
		x := int((point[0]-minX)*scaleX + float64((paddingX+1)*c.RuneSize().Width))
		y := int((point[1]-minY)*scaleY + float64(paddingY*c.RuneSize().Height))

		if first {
			first = false
			prevX = x
			prevY = y
		}

		if prevX <= x {
			c.DrawLine(prevY, prevX, y, x)
		}

		prevX = x
		prevY = y
	}
	c.drawAxes(paddingX, paddingY, minX, maxX, minY, maxY)

	b := bytes.NewBuffer(nil)
	c.GetBuffer().Render(b)
	return b.String()
}

func roundDownToPercentOfRange(x, d float64) float64 {
	return math.Floor((x*(100-math.Copysign(5, x)))/d) * d / 100
}

func roundUpToPercentOfRange(x, d float64) float64 {
	return math.Ceil((x*105)/d) * d / 100
}

func minMax(table *data.Table) (minX, maxX, minY, maxY float64) {
	minX, minY = math.Inf(1), math.Inf(1)
	maxX, maxY = math.Inf(-1), math.Inf(-1)

	for _, r := range table.Rows {
		if len(r) < 2 {
			continue
		}
		maxX = math.Max(maxX, r[0])
		minX = math.Min(minX, r[0])
		maxY = math.Max(maxY, r[1])
		minY = math.Min(minY, r[1])
	}

	yRange := maxY - minY
	minY = roundDownToPercentOfRange(minY, yRange)
	maxY = roundUpToPercentOfRange(maxY, yRange)
	return
}
