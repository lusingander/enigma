package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lusingander/enigma/enigma"
	"github.com/muesli/reflow/wrap"
)

var (
	titleStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderTop(false).
			BorderRight(false).
			BorderLeft(false).
			BorderForeground(lipgloss.Color("240")).
			Foreground(lipgloss.Color("68")).
			Bold(true).
			Align(lipgloss.Center)

	contentStyle = lipgloss.NewStyle().
			Padding(0, 2)

	labelStyle = lipgloss.NewStyle().
			Padding(0, 3).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("68")).
			Foreground(lipgloss.Color("68"))

	rotorRowStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Margin(0, 1).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderTop(false).
			BorderBottom(false).
			Foreground(lipgloss.Color("240"))

	rotorRowSelectedStyle = rotorRowStyle.Copy().
				Foreground(lipgloss.Color("252"))

	lampboardStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Foreground(lipgloss.Color("240"))

	lampboardSelectedStyle = lampboardStyle.Copy().
				BorderForeground(lipgloss.Color("214")).
				Foreground(lipgloss.Color("214"))

	keyboardStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("108"))

	keyboardSelectedStyle = keyboardStyle.Copy().
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("226")).
				Foreground(lipgloss.Color("226"))

	plugboardStyle = lipgloss.NewStyle().
			Padding(0, 2)

	rowCenterStyle = lipgloss.NewStyle().
			Padding(0, 2)
)

var (
	keysTop    = []string{"Q", "W", "E", "R", "T", "Z", "U", "I", "O"}
	keysCenter = []string{"A", "S", "D", "F", "G", "H", "J", "K"}
	keysBottom = []string{"P", "Y", "X", "C", "V", "B", "N", "M", "L"}

	plugTop    = []string{"￭", "￭", "￭", "￭", "￭", "￭", "￭", "￭", "￭"}
	plugCenter = []string{"￭", "￭", "￭", "￭", "￭", "￭", "￭", "￭"}
	plugBottom = []string{"￭", "￭", "￭", "￭", "￭", "￭", "￭", "￭", "￭"}
)

type model struct {
	enigma.Enigma

	openSettingsView bool

	width, height int

	inputViewport  viewport.Model
	outputViewport viewport.Model
	input, output  string
}

var _ tea.Model = (*model)(nil)

func newModel() model {
	m := model{
		openSettingsView: false,
		inputViewport:    viewport.New(0, 0),
		outputViewport:   viewport.New(0, 0),
	}
	m.reset()
	return m
}

func (m *model) setSize(w, h int) {
	m.width = w
	m.height = h

	// todo: fix
	vw := m.width - 60
	m.inputViewport.Width = vw
	m.outputViewport.Width = vw
	vh := (m.height - 12) / 2
	m.inputViewport.Height = vh
	m.outputViewport.Height = vh
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch s := msg.String(); s {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.reset()
		case "tab":
			m.openSettingsView = !m.openSettingsView
		default:
			m.handleKeyInput(s)
		}
	case tea.WindowSizeMsg:
		m.setSize(msg.Width, msg.Height)
	}
	return m, nil
}

func (m *model) reset() {
	// todo: fix
	m.Enigma = enigma.NewM3(
		[3]enigma.Rotor{
			enigma.NewRotor1('A', 'A'),
			enigma.NewRotor2('A', 'A'),
			enigma.NewRotor3('A', 'A'),
		},
		enigma.NewReflectorB(),
		enigma.NewPlugboard(),
	)
	m.input = ""
	m.output = ""
	m.inputViewport.SetContent("")
	m.outputViewport.SetContent("")
}

func (m *model) handleKeyInput(key string) {
	if !isOneLetterAlpha(key) {
		return
	}
	if m.openSettingsView {
		// todo: fix
		return
	}
	m.send(key)
	return
}

func (m *model) send(key string) {
	r := toUpperRune(key)
	encoded := m.Enigma.EncodeRune(r)
	m.updateViewportContent(r, encoded)
}

func (m *model) updateViewportContent(key, encoded rune) {
	m.input += string(key)
	m.output += string(encoded)
	input := wrap.String(m.input, m.inputViewport.Width)
	output := wrap.String(m.output, m.outputViewport.Width)
	m.inputViewport.SetContent(input)
	m.outputViewport.SetContent(output)
}

func (m model) View() string {
	var content string
	if m.openSettingsView {
		content = m.settinsView()
	} else {
		content = m.simulatorView()
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.titleView(),
		content,
	)
}

func (m model) titleView() string {
	return titleStyle.Width(m.width).Render("ENIGMA SIMULATOR")
}

func (m model) settinsView() string {
	// todo: impl
	return ""
}

func (m model) simulatorView() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		contentStyle.Render(m.enigmaView()),
		contentStyle.Render(m.inOutView()),
	)
}

func (m model) inOutView() string {
	return lipgloss.NewStyle().Padding(0, 0, 0, 5).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.inputView(),
			m.outputView(),
		),
	)
}

func (m model) inputView() string {
	label := labelStyle.Render("INPUT")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		m.inputViewport.View(),
	)
}

func (m model) outputView() string {
	label := labelStyle.Render("OUTPUT")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		m.outputViewport.View(),
	)
}

func (m model) enigmaView() string {
	rotors := m.rotorsView()
	lampboard := m.lampboardView()
	keyboard := m.keyboardView()
	// plugboard := m.plugboardView()
	return lipgloss.JoinVertical(
		lipgloss.Left,
		rotors,
		lampboard,
		keyboard,
		// plugboard,
	)
}

func (m model) rotorsView() string {
	label := labelStyle.Render("ROTORS")
	rps := m.RotorPositions()
	rotorUpperRunes := make([]string, len(rps))
	rotorSelectedRunes := make([]string, len(rps))
	rotorLowerRunes := make([]string, len(rps))
	for i, rp := range rps {
		// right to left
		i = len(rps) - 1 - i
		rotorUpperRunes[i] = rotorRowStyle.Render(string(prevAlpha(rp)))
		rotorSelectedRunes[i] = rotorRowSelectedStyle.Render(string(rp))
		rotorLowerRunes[i] = rotorRowStyle.Render(string(nextAlpha(rp)))
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		lipgloss.JoinHorizontal(lipgloss.Bottom, rotorUpperRunes...),
		lipgloss.JoinHorizontal(lipgloss.Bottom, rotorSelectedRunes...),
		lipgloss.JoinHorizontal(lipgloss.Bottom, rotorLowerRunes...),
	)
}

func (m model) lampboardView() string {
	label := labelStyle.Render("LAMPBOARD")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		m.lampboardRowView(keysTop),
		rowCenterStyle.Render(m.lampboardRowView(keysCenter)),
		m.lampboardRowView(keysBottom),
	)
}

func (m model) lampboardRowView(ks []string) string {
	return m.rowView(lampboardStyle, lampboardSelectedStyle, ks, lastRune(m.output))
}

func (m model) keyboardView() string {
	label := labelStyle.Render("KEYBOARD")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		m.keyboardRowView(keysTop),
		rowCenterStyle.Render(m.keyboardRowView(keysCenter)),
		m.keyboardRowView(keysBottom),
	)
}

func (m model) keyboardRowView(ks []string) string {
	return m.rowView(keyboardStyle, keyboardSelectedStyle, ks, lastRune(m.input))
}

func (m model) plugboardView() string {
	label := labelStyle.Render("PLUGBOARD")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		m.plugboardRowView(keysTop),
		m.plugboardRowView(plugTop),
		rowCenterStyle.Render(m.plugboardRowView(keysCenter)),
		rowCenterStyle.Render(m.plugboardRowView(plugCenter)),
		m.plugboardRowView(keysBottom),
		m.plugboardRowView(plugBottom),
	)
}

func (m model) plugboardRowView(ks []string) string {
	return m.rowView(plugboardStyle, plugboardStyle, ks, '0')
}

func (m model) rowView(normalStyle, selectedStyle lipgloss.Style, ks []string, selected rune) string {
	ss := make([]string, 0, len(ks))
	for _, k := range ks {
		var s string
		if k == string(selected) {
			s = selectedStyle.Render(k)
		} else {
			s = normalStyle.Render(k)
		}
		ss = append(ss, s)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, ss...)
}

func Start() error {
	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	return p.Start()
}
