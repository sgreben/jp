package draw

import (
	"io"
)

type Buffer struct {
	Runes [][]rune
	Box
}

func (b *Buffer) GetBuffer() *Buffer { return b }

func NewBuffer(size Box) *Buffer {
	r := make([][]rune, size.Height)
	for i := range r {
		r[i] = make([]rune, size.Width)
	}
	return &Buffer{
		Runes: r,
		Box:   size,
	}
}

func (b *Buffer) Fill(r rune) {
	for i := range b.Runes {
		for j := range b.Runes[i] {
			b.Runes[i][j] = r
		}
	}
}

func (b *Buffer) Get(y, x int) rune {
	if y < 0 || y >= b.Height {
		return 0
	}
	if x < 0 || x >= b.Width {
		return 0
	}
	return b.Runes[y][x]
}

func (b *Buffer) Set(y, x int, r rune) {
	if y < 0 || y >= b.Height {
		return
	}
	if x < 0 || x >= b.Width {
		return
	}
	b.Runes[y][x] = r
}

func (b *Buffer) SetRow(y, x1, x2 int, r rune) {
	for i := x1; i < x2; i++ {
		b.Set(y, i, r)
	}
}

func (b *Buffer) SetColumn(y1, y2, x int, r rune) {
	for i := y1; i < y2; i++ {
		b.Set(i, x, r)
	}
}

func (b *Buffer) SetOr(y, x int, r rune) {
	if y < 0 || y >= b.Height {
		return
	}
	if x < 0 || x >= b.Width {
		return
	}
	b.Runes[y][x] |= r
}

func (b *Buffer) Write(y, x int, r []rune) {
	if y < 0 || y >= b.Height {
		return
	}
	for i := 0; i < len(r); i++ {
		xi := x + i
		if xi < 0 || xi >= b.Width {
			continue
		}
		b.Runes[y][xi] = r[i]
	}
}

func (b *Buffer) WriteLeft(y, x int, r []rune) {
	b.Write(y, x-len(r), r)
}

func (b *Buffer) WriteRight(y, x int, r []rune) {
	b.Write(y, x, r)
}

func (b *Buffer) WriteCenter(y, x int, r []rune) {
	b.Write(y, x-len(r)/2, r)
}

func (b *Buffer) Render(w io.Writer) {
	for i := range b.Runes {
		row := b.Runes[b.Height-i-1]
		w.Write([]byte(string(row)))
		if i < b.Height-1 {
			w.Write([]byte{'\n'})
		}
	}
}
