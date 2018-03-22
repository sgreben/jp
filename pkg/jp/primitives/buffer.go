package primitives

func Buffer(size int) (out []string) {
	out = make([]string, size)
	for i := range out {
		out[i] = " "
	}
	return
}
