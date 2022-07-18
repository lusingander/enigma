package enigma

type idx int

type idxMap map[idx]idx

func buildMap(w string) idxMap {
	m := make(idxMap)
	for i, r := range w {
		m[idx(i)] = toIdx(r)
	}
	return m
}

func toIdx(r rune) idx {
	return idx(r - 'A')
}

func toRune(i idx) rune {
	return rune(i + 'A')
}

func mod26(i idx) idx {
	return ((i % 26) + 26) % 26
}
