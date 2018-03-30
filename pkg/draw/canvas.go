package draw

type Pixels interface {
	GetBuffer() *Buffer
	Clear()
	Set(y, x int)
	Size() Box
}

type Canvas struct{ Pixels }

func (c *Canvas) RuneSize() Box {
	pixelBox := c.Size()
	runeBox := c.GetBuffer().Box
	return Box{
		Width:  pixelBox.Width / runeBox.Width,
		Height: pixelBox.Height / runeBox.Height,
	}
}

func (c *Canvas) DrawLine(y0, x0, y1, x1 int) {
	line(y0, x0, y1, x1, c.Set)
}
