package draw

var _ Pixels = &Full{}

const fullBlock = '█'

type Full struct{ *Buffer }

func (b *Full) Size() Box { return b.Box }

func (b *Full) Set(y, x int) { b.Buffer.Set(y, x, fullBlock) }

func (b *Full) Clear() { b.Fill(' ') }
