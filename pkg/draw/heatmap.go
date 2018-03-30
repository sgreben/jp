package draw

type Heatmap struct{ *Buffer }

func (b *Heatmap) Size() Box { return b.Box }

var shades = []rune(" ·░▒▒▒▒▓▓▓▓█")
var nonZeroShade = 1.0 / float64(len(shades)-1)

func (b *Heatmap) Set(y, x int, fill float64) {
	if fill < 0.0 {
		fill = 0.0
	}
	if fill > 0.0 && fill < nonZeroShade {
		fill = nonZeroShade
	}
	if fill > 1.0 {
		fill = 1.0
	}
	b.Buffer.Set(y, x, shades[int(fill*float64(len(shades)-1))])
}

func (b *Heatmap) Clear() { b.Fill(shades[0]) }
