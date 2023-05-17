package thespine

const THE_SIZE = 3

func Decode(s string) string {
	sr := []rune(s)
	l := len(sr)
	g := make([][]rune, 0)
	gc := l / THE_SIZE
	if l%THE_SIZE != 0 {
		gc++
	}
	for i := 0; i < gc; i++ {
		si := l - (i+1)*THE_SIZE
		ei := l - i*THE_SIZE
		if si < 0 {
			si = 0
		}
		gs := sr[si:ei]
		g = append(g, gs)
	}
	return concat(g)
}

func Encode(s string) string {
	sr := []rune(s)
	l := len(sr)
	g := make([][]rune, 0)
	gc := l / THE_SIZE
	if l%THE_SIZE != 0 {
		gc++
	}
	for i := 0; i < gc; i++ {
		si := i * THE_SIZE
		ei := (i + 1) * THE_SIZE
		if ei > l {
			ei = l
		}
		gs := sr[si:ei]
		g = append(g, gs)
	}
	for i, j := 0, len(g)-1; i < j; i, j = i+1, j-1 {
		g[i], g[j] = g[j], g[i]
	}
	return concat(g)
}

func concat(r [][]rune) string {
	var s string
	for _, r := range r {
		s += string(r)
	}
	return s
}
