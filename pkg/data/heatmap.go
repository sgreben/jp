package data

import "math"

type Heatmap struct {
	X, Y       []Bin
	Z          [][]float64
	MinX, MaxX uint64
	MinY, MaxY uint64
	MinZ, MaxZ uint64
}

func NewHeatmap(x, y []Bin, z [][]uint64) *Heatmap {
	h := new(Heatmap)
	h.X, h.Y = x, y
	h.Z = make([][]float64, len(z))
	h.MinX, h.MinY, h.MinZ = math.MaxUint64, math.MaxUint64, math.MaxUint64
	h.MaxX, h.MaxY, h.MaxZ = 0, 0, 0
	for _, b := range x {
		if b.Count > h.MaxX {
			h.MaxX = b.Count
		}
		if b.Count < h.MinX {
			h.MinX = b.Count
		}
	}
	for _, b := range x {
		b.CountNorm = float64(b.Count-h.MinX) / float64(h.MaxX-h.MinX)
	}
	for _, b := range y {
		if b.Count > h.MaxY {
			h.MaxY = b.Count
		}
		if b.Count < h.MinY {
			h.MinY = b.Count
		}
	}
	for _, b := range y {
		b.CountNorm = float64(b.Count-h.MinY) / float64(h.MaxY-h.MinY)
	}
	for i := range z {
		h.Z[i] = make([]float64, len(z[i]))
		for _, b := range z[i] {
			if b > h.MaxZ {
				h.MaxZ = b
			}
			if b < h.MinZ {
				h.MinZ = b
			}
		}
	}
	for i := range z {
		for j := range z[i] {
			h.Z[i][j] = float64(z[i][j]-h.MinZ) / float64(h.MaxZ-h.MinZ)
		}
	}
	if h.MaxX == 0 {
		h.MaxX = 1
	}
	if h.MaxY == 0 {
		h.MaxY = 1
	}
	if h.MaxZ == 0 {
		h.MaxZ = 1
	}

	return h
}
