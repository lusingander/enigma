package enigma

import (
	"strings"
	"testing"
)

func TestM3(t *testing.T) {
	tests := []struct {
		rotors    [3]Rotor
		reflector Reflector
		plugboard Plugboard
		s         string
		want      string
	}{
		{
			rotors: [3]Rotor{
				NewRotor1('A', 'A'),
				NewRotor2('A', 'A'),
				NewRotor3('A', 'A'),
			},
			reflector: NewReflectorB(),
			plugboard: NewPlugboard(),
			s:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:      "FUVEPUMWARVQKEFGHGDIJFMFXI",
		},
		{
			rotors: [3]Rotor{
				NewRotor1('C', 'A'),
				NewRotor2('D', 'A'),
				NewRotor3('E', 'A'),
			},
			reflector: NewReflectorB(),
			plugboard: NewPlugboard(),
			s:         "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:      "MRMQZKWEZCGBGGTZUOWQIOOHMOTULCKLVCXQDZBJQRSFTZVAJAWQFJSFLKYENWNQNDEZSMXLRMPTWQTUEVQAVDPHSRAZQUPFXRREFMDAWXYEABUISPCRADPELHNWPKRQLRYAROGWJJUEIEWIGZISDZXCLFDATUSILPVIWDFNPRWTBBQNSHDVRJMRQHZHSFZCGRASKJFBHRALEWCDHIYCKHNSAHVMOJSOXZNVAWGYGVOEBTQGIDMOYWXYTCIBAYSCKFIS",
		},
		{
			rotors: [3]Rotor{
				NewRotor6('F', 'A'),
				NewRotor4('T', 'A'),
				NewRotor5('V', 'A'),
			},
			reflector: NewReflectorB(),
			plugboard: NewPlugboard(),
			s:         "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:      "YTDUFSPJYVMJPHIQPGQMDGUCPEOFFAHXSOXKHBCOQFSQHSPMMOQWGEJVNHJYRXRAGDZGGYTEJZBHVTFLGNKBROYGGFHCPBZQCOSDATJKPOLJQCBIVIFYWLTNHSEXMBXASECDXMVEFDORXQAYVCFXPJTMXOQYGDJZULJIVUSWYRJQHSFRGZIZKCUVRVLCVQTGOYHWPSNOQMBAFSTYYKSBKQMBXYWQRQMZHTCXYRDFCYEHZLCMXGYWTMYFGVEWYYCGTPOUKCKAJAQGSRHWNOVQDBYQVGBUUKBLGAWKTKNOMTLHYTDOXZYFFUHVBKZKPDZRJIEJIZTAXCMBNHLVVRTZRCJBVDPWDMAKGILMXSXJHAGNCWNYIEENWTHPFOWFKIEVTFQDBQVAFLLLXZTHPPRBSCNAUYQCAFLPEOEJJUIWQGLNHAPSTUJCAGJLZSRTBAKTDLXVIPXGIMLZQMVZRJPKDNIXQKSJQSSIGRXYROOEAPKRRSMMIPSGWAUXWGUQZXKUNJGGVSLXYGNAXDZWMFDBWKDSJMIPOCENXHWISZURDDPGXNHYPZFTCYDJPZDTUUVYNHRQGKMFPHQNZEBAXEPUGPBUTRSUPULMMIQHLUAWOHGPLARYCZMPGEUXSNGCUCBAWZTFJAQRQHMAJNNHMIBMZPEMLCMHFUZZJDZQVTPPYTYJKXIJPGZQHKVWZUCVTETRBIKWPXXTSFWRWSZGMSMIHENEQWFPFCEXLNTUSFBAWZMGCPIPRLKVSLXNXVNBTXLFCALUAZKPIRRK",
		},
		{
			rotors: [3]Rotor{
				NewRotor2('E', 'C'),
				NewRotor3('A', 'F'),
				NewRotor1('Z', 'Z'),
			},
			reflector: NewReflectorC(),
			plugboard: NewPlugboard(),
			s:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:      "CRPYXXUUYZOCYRQNBLERAHHODJ",
		},
		{
			rotors: [3]Rotor{
				NewRotor6('X', 'D'),
				NewRotor7('N', 'J'),
				NewRotor8('P', 'O'),
			},
			reflector: NewReflectorC(),
			plugboard: NewPlugboard(),
			s:         "AAAAAAAAAABBBBBBBBBBCCCCCCCCCCDDDDDDDDDDEEEEEEEEEEFFFFFFFFFFGGGGGGGGGGHHHHHHHHHHIIIIIIIIIIJJJJJJJJJJKKKKKKKKKKLLLLLLLLLLMMMMMMMMMMNNNNNNNNNNOOOOOOOOOOPPPPPPPPPPQQQQQQQQQQRRRRRRRRRRSSSSSSSSSSTTTTTTTTTTUUUUUUUUUUVVVVVVVVVVWWWWWWWWWWXXXXXXXXXXYYYYYYYYYYZZZZZZZZZZ",
			want:      "ISBQFMSYHNVWWAECXUGIMKKTZWRWZWMQAECQXCFJUFWGJLSRDLKLLLUNLOGHDNJIHAIHYQFAMDZLZPYOBNNLJZTRUTEGMVCIRHNEAJHLHMUIFXHTOCNBMWRQWNFDDPWUACDQRKEACLXHUSTWBDFPZAGHFAUEQEQZTBGCZNDMYDVBNIYCMJOXERRUTMRKZNZOPSCLURIBEZVNSABNDSMJFRCIFWXYJXPYHATRLOWLFDQIGICTIKLVZLWBUKMWIUHPJTRL",
		},
		{
			rotors: [3]Rotor{
				NewRotor1('A', 'A'),
				NewRotor2('A', 'A'),
				NewRotor3('A', 'A'),
			},
			reflector: NewReflectorB(),
			plugboard: NewPlugboard(
				Pair{'C', 'H'},
				Pair{'E', 'Z'},
				Pair{'Q', 'N'},
			),
			s:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want: "FUWZOUMQARVNKRFGBGDIJFMFXJ",
		},
	}

	for _, test := range tests {
		e := NewM3(test.rotors, test.reflector, test.plugboard)
		got := e.EncodeString(test.s)
		if got != test.want {
			t.Errorf("got=%s, want=%s", got, test.want)
		}
	}
}

func TestM3_long(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	rs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := strings.Repeat(rs, 10000)

	newM3 := func() Enigma {
		return NewM3(
			[3]Rotor{
				NewRotor1('A', 'A'),
				NewRotor2('A', 'A'),
				NewRotor3('A', 'A'),
			},
			NewReflectorB(),
			NewPlugboard(),
		)
	}
	e1 := newM3()
	e2 := newM3()

	encoded := e1.EncodeString(s)
	decoded := e2.EncodeString(encoded)

	if s != decoded {
		t.Errorf("s=%s, encoded=%s, decoded=%s", s, encoded, decoded)
	}
}

func TestM3_pos(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	rs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := strings.Repeat(rs, 10)

	for _, r1 := range rs {
		for _, r2 := range rs {
			for _, r3 := range rs {
				testPos(r1, r2, r3, s, t)
			}
		}
	}
}

func testPos(r1, r2, r3 rune, s string, t *testing.T) {
	newM3 := func() Enigma {
		return NewM3(
			[3]Rotor{
				NewRotor1(r1, 'A'),
				NewRotor2(r2, 'A'),
				NewRotor3(r3, 'A'),
			},
			NewReflectorB(),
			NewPlugboard(),
		)
	}
	e1 := newM3()
	e2 := newM3()

	encoded := e1.EncodeString(s)
	decoded := e2.EncodeString(encoded)

	if s != decoded {
		t.Errorf("s=%s, encoded=%s, decoded=%s, pos=(%s, %s, %s)",
			s, encoded, decoded, string(r1), string(r2), string(r3))
	}
}

func TestM3_ring(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	rs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := strings.Repeat(rs, 10)

	for _, r1 := range rs {
		for _, r2 := range rs {
			for _, r3 := range rs {
				testRing(r1, r2, r3, s, t)
			}
		}
	}
}

func testRing(r1, r2, r3 rune, s string, t *testing.T) {
	newM3 := func() Enigma {
		return NewM3(
			[3]Rotor{
				NewRotor1('A', r1),
				NewRotor2('A', r2),
				NewRotor3('A', r3),
			},
			NewReflectorB(),
			NewPlugboard(),
		)
	}
	e1 := newM3()
	e2 := newM3()

	encoded := e1.EncodeString(s)
	decoded := e2.EncodeString(encoded)

	if s != decoded {
		t.Errorf("s=%s, encoded=%s, decoded=%s, ring=(%s, %s, %s)",
			s, encoded, decoded, string(r1), string(r2), string(r3))
	}
}
