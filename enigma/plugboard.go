package enigma

// Plugboard (Steckerbrett)
type Plugboard interface {
	substitute(idx) idx
}

type Pair struct {
	Fst, Snd rune
}

func NewPlugboard(pairs ...Pair) Plugboard {
	m := make(idxMap)
	for i := 0; i < 26; i++ {
		idx := idx(i)
		m[idx] = idx
	}
	for _, p := range pairs {
		i1 := toIdx(p.Fst)
		i2 := toIdx(p.Snd)
		m[i1], m[i2] = m[i2], m[i1]
	}
	return &plugboard{
		idxMap: m,
	}
}

type plugboard struct {
	idxMap
}

func (s *plugboard) substitute(i idx) idx {
	return s.idxMap[i]
}
