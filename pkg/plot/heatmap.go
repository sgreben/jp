package plot

import (
	"bytes"

	"github.com/sgreben/jp/pkg/data"
	"github.com/sgreben/jp/pkg/draw"
)

// HeatMap is a heatmap
type HeatMap struct{ draw.Heatmap }

// NewHeatMap returns a new line chart
func NewHeatMap(buffer *draw.Buffer) *HeatMap { return &HeatMap{draw.Heatmap{buffer}} }

func (c *HeatMap) drawAxes(paddingX, paddingY int, minX, maxX, minY, maxY float64) {
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
func (c *HeatMap) Draw(heatmap *data.Heatmap) string {
	var scaleY, scaleX float64

	minX := heatmap.X[0].LeftInclusive
	maxX := heatmap.X[len(heatmap.X)-1].Right
	minY := heatmap.Y[0].LeftInclusive
	maxY := heatmap.Y[len(heatmap.Y)-1].Right
	minLabelWidth := len(Ff(minY))
	maxLabelWidth := len(Ff(maxY))

	paddingX := minLabelWidth + 1
	paddingY := 2
	if minLabelWidth < maxLabelWidth {
		paddingX = maxLabelWidth + 1
	}
	chartWidth := c.Size().Width - (paddingX + 1)
	chartHeight := c.Size().Height - paddingY
	scaleX = float64(chartWidth) / (maxX - minX)
	scaleY = float64(chartHeight) / (maxY - minY)

	for i := range heatmap.Z {
		for j := range heatmap.Z[i] {
			x0 := int((heatmap.X[j].LeftInclusive-minX)*scaleX + float64(paddingX+1))
			y0 := int((heatmap.Y[i].LeftInclusive-minY)*scaleY + float64(paddingY))
			x1 := int((heatmap.X[j].Right-minX)*scaleX + float64(paddingX+1))
			y1 := int((heatmap.Y[i].Right-minY)*scaleY + float64(paddingY))
			z := heatmap.Z[i][j]
			for x := x0; x < x1; x++ {
				for y := y0; y < y1; y++ {
					c.Set(y, x, z)
				}
			}
		}
	}
	c.drawAxes(paddingX, paddingY, minX, maxX, minY, maxY)

	b := bytes.NewBuffer(nil)
	c.GetBuffer().Render(b)
	return b.String()
}
