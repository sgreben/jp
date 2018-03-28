package plot

import (
	"bytes"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
)

// ScatterChart is a scatter plot
type ScatterChart struct{ draw.Canvas }

// NewScatterChart returns a new line chart
func NewScatterChart(canvas draw.Canvas) *ScatterChart { return &ScatterChart{canvas} }

func (c *ScatterChart) drawAxes(paddingX, paddingY int, minX, maxX, minY, maxY float64) {
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
func (c *ScatterChart) Draw(table *data.Table) string {
	var scaleY, scaleX float64

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

	for _, point := range table.Rows {
		if len(point) < 2 {
			continue
		}
		x := int((point[0]-minX)*scaleX + float64((paddingX+1)*c.RuneSize().Width))
		y := int((point[1]-minY)*scaleY + float64(paddingY*c.RuneSize().Height))
		c.Set(y, x)
	}
	c.drawAxes(paddingX, paddingY, minX, maxX, minY, maxY)

	b := bytes.NewBuffer(nil)
	c.GetBuffer().Render(b)
	return b.String()
}
