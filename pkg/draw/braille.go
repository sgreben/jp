package draw

var _ Pixels = &Braille{}

const (
	brailleScaleX = 2
	brailleScaleY = 4
	brailleEmpty  = rune(0x2800)
)

type Braille struct{ *Buffer }

func (b *Braille) Size() Box {
	return Box{
		Width:  b.Width * brailleScaleX,
		Height: b.Height * brailleScaleY,
	}
}

func (b *Braille) Clear() { b.Fill(brailleEmpty) }

func (b *Braille) Set(y, x int) {
	ry := y / brailleScaleY
	rx := x / brailleScaleX
	b.SetOr(ry, rx, braillePoint(y, x))
}

func braillePoint(y, x int) rune {
	var cy, cx int
	if y >= 0 {
		cy = y % 4
	} else {
		cy = 3 + ((y + 1) % 4)
	}
	if x >= 0 {
		cx = x % 2
	} else {
		cx = 1 + ((x + 1) % 2)
	}
	pixelMap := [4][2]rune{{1, 8}, {2, 16}, {4, 32}, {64, 128}}
	return pixelMap[3-cy][cx]
}
