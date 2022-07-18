package enigma

const (
	reflectorB_Wiring = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
	reflectorC_Wiring = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
)

// Reflector (Umkehrwalze)
type Reflector interface {
	substitute(idx) idx
}

func NewReflectorB() Reflector {
	return newReflector(reflectorB_Wiring)
}

func NewReflectorC() Reflector {
	return newReflector(reflectorC_Wiring)
}

type reflector struct {
	idxMap
}

func newReflector(w string) Reflector {
	return &reflector{
		idxMap: buildMap(w),
	}
}

func (u reflector) substitute(i idx) idx {
	return u.idxMap[i]
}
