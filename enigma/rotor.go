package enigma

const (
	rotor1_Wiring = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor2_Wiring = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	rotor3_Wiring = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	rotor4_Wiring = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
	rotor5_Wiring = "VZBRGITYUPSDNHLXAWMJQOFECK"
	rotor6_Wiring = "JPGVOUMFYQBENHZRDKASXLICTW"
	rotor7_Wiring = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
	rotor8_Wiring = "FKQHTLXOCBJSPDZRAMEWNIUYGV"

	rotor1_Notch = "Q"
	rotor2_Notch = "E"
	rotor3_Notch = "V"
	rotor4_Notch = "J"
	rotor5_Notch = "Z"
	rotor6_Notch = "ZM"
	rotor7_Notch = "ZM"
	rotor8_Notch = "ZM"
)

type Rotor interface {
	substituteRtoL(idx) idx
	substituteLtoR(idx) idx
	turnover() bool
	step()
}

func NewRotor1(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor1_Wiring, rotor1_Notch)
}

func NewRotor2(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor2_Wiring, rotor2_Notch)
}

func NewRotor3(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor3_Wiring, rotor3_Notch)
}

func NewRotor4(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor4_Wiring, rotor4_Notch)
}

func NewRotor5(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor5_Wiring, rotor5_Notch)
}

func NewRotor6(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor6_Wiring, rotor6_Notch)
}

func NewRotor7(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor7_Wiring, rotor7_Notch)
}

func NewRotor8(pos, ring rune) Rotor {
	return newRotor(pos, ring, rotor8_Wiring, rotor8_Notch)
}

type rotor struct {
	rIdxMap idxMap
	lIdxMap idxMap
	pos     idx
	ring    idx
	notchSet

	wiring string
}

func newRotor(pos, ring rune, wiring, notch string) Rotor {
	rotor := &rotor{
		rIdxMap:  make(idxMap),
		lIdxMap:  make(idxMap),
		pos:      toIdx(pos),
		ring:     toIdx(ring),
		notchSet: newNotchSet(notch),
		wiring:   wiring,
	}
	rotor.updateIdxMap()
	return rotor
}

func (w *rotor) updateIdxMap() {
	for i, r := range w.wiring {
		idx := mod26(idx(i) - w.pos + w.ring)
		w.rIdxMap[idx] = mod26(toIdx(r) - w.pos + w.ring)
	}
	for k, v := range w.rIdxMap {
		w.lIdxMap[v] = k
	}
}

func (w rotor) substituteRtoL(i idx) idx {
	return w.rIdxMap[i]
}

func (w rotor) substituteLtoR(i idx) idx {
	return w.lIdxMap[i]
}

func (w rotor) turnover() bool {
	_, ok := w.notchSet[w.pos]
	return ok
}

func (w *rotor) step() {
	w.pos = mod26(w.pos + 1)
	w.updateIdxMap()
}

type notchSet map[idx]any

func newNotchSet(n string) notchSet {
	s := make(notchSet)
	for _, r := range n {
		i := toIdx(r)
		s[i] = struct{}{}
	}
	return s
}
