package draw

var _ Pixels = &Quarter{}

const (
	quarterScaleX = 2
	quarterScaleY = 2
)

type Quarter struct{ *Buffer }

func (b *Quarter) Size() Box {
	return Box{
		Width:  b.Width * quarterScaleX,
		Height: b.Height * quarterScaleY,
	}
}

func (b *Quarter) Set(y, x int) {
	ry, cy := y/2, 1-y%2
	rx, cx := x/2, x%2
	i := index(b.Get(ry, rx))
	b.Buffer.Set(ry, rx, runes[i|(1<<uint(cx+2*cy))])
}

func (b *Quarter) Clear() { b.Fill(runes[0]) }

var runes = []rune(" ▘▝▀▖▌▞▛▗▚▐▜▄▙▟█")

func index(r rune) int {
	for i := range runes {
		if runes[i] == r {
			return i
		}
	}
	return 0
}
