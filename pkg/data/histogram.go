package data

import (
	"fmt"
	"math"
)

import "strconv"

const maxDigits = 6

func ff(x float64) string {
	minExact := strconv.FormatFloat(x, 'g', -1, 64)
	fixed := strconv.FormatFloat(x, 'g', maxDigits, 64)
	if len(minExact) < len(fixed) {
		return minExact
	}
	return fixed
}

type Bin struct {
	LeftInclusive  float64
	Right          float64
	RightInclusive bool
	Count          uint64
}

func (b *Bin) String() string {
	if b.RightInclusive {
		return fmt.Sprintf("[%s,%s]", ff(b.LeftInclusive), ff(b.Right))
	}
	return fmt.Sprintf("[%s,%s)", ff(b.LeftInclusive), ff(b.Right))
}

type Bins struct {
	Number    int
	min, max  float64
	numPoints int
}

func (b *Bins) ChooseSqrt() {
	b.Number = int(math.Sqrt(float64(b.numPoints)))
}

func (b *Bins) ChooseSturges() {
	b.Number = int(math.Ceil(math.Log2(float64(b.numPoints))) + 1)
}

func (b *Bins) ChooseRice() {
	b.Number = int(math.Ceil(2 * math.Pow(float64(b.numPoints), 1.0/3.0)))
}

func NewBins(points []float64) *Bins {
	bins := new(Bins)
	bins.numPoints = len(points)
	bins.Number = 5
	bins.min = math.Inf(1)
	bins.max = math.Inf(-1)
	for _, x := range points {
		bins.min = math.Min(bins.min, x)
		bins.max = math.Max(bins.max, x)
	}
	return bins
}

func (b *Bins) left(i int) float64 {
	return b.min + ((b.max - b.min) / float64(b.Number) * float64(i))
}

func (b *Bins) right(i int) float64 {
	return b.left(i + 1)
}

func (b *Bins) All() (out []Bin) {
	if b.max == b.min {
		b.Number = 1
	}
	for i := 0; i < b.Number; i++ {
		out = append(out, Bin{
			LeftInclusive: b.left(i),
			Right:         b.right(i),
		})
	}
	out[b.Number-1].RightInclusive = true
	return
}

func (b *Bins) Point(x float64) int {
	if b.max == b.min {
		return 0
	}
	i := int((x - b.min) / (b.max - b.min) * float64(b.Number))
	if i >= b.Number {
		i--
	}
	return i
}

func Histogram(points []float64, bins *Bins) (out []Bin) {
	out = bins.All()
	for _, b := range out {
		b.Count = 0
	}
	for _, x := range points {
		out[bins.Point(x)].Count++
	}
	return
}
