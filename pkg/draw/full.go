package draw

var _ Pixels = &Full{}

type Full struct{ *Buffer }

func (b *Full) Size() Box { return b.Box }

func (b *Full) Set(y, x int) { b.Buffer.Set(y, x, '█') }

func (b *Full) Clear() { b.Fill(' ') }
