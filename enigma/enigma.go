package enigma

type Enigma interface {
	EncodeRune(rune) rune
	EncodeString(string) string
}

func NewM3(rotors [3]Rotor, reflector Reflector, plugboard Plugboard) Enigma {
	return &M3{
		rotors:    rotors,
		reflector: reflector,
		plugboard: plugboard,
	}
}

type M3 struct {
	rotors    [3]Rotor
	reflector Reflector
	plugboard Plugboard
}

func (e *M3) EncodeRune(r rune) rune {
	if e.rotors[1].turnover() {
		e.rotors[1].step()
		e.rotors[2].step()
	}
	if e.rotors[0].turnover() {
		e.rotors[1].step()
	}
	e.rotors[0].step()

	i := toIdx(r)

	i = e.plugboard.substitute(i)
	i = e.rotors[0].substituteRtoL(i)
	i = e.rotors[1].substituteRtoL(i)
	i = e.rotors[2].substituteRtoL(i)
	i = e.reflector.substitute(i)
	i = e.rotors[2].substituteLtoR(i)
	i = e.rotors[1].substituteLtoR(i)
	i = e.rotors[0].substituteLtoR(i)
	i = e.plugboard.substitute(i)

	return toRune(i)
}

func (e *M3) EncodeString(s string) string {
	rs := make([]rune, len(s))
	for i, r := range s {
		rs[i] = e.EncodeRune(r)
	}
	return string(rs)
}
