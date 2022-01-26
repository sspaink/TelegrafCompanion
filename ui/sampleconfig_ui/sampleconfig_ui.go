package sampleconfig_ui

import tea "github.com/charmbracelet/bubbletea"

const (
	welcomePage = iota
	pluginSelection
)

var (
	ready bool
)

// A global value because SampleConfigUI update and view are pass by value to adhere to tea.Model interface
// This is to allow other page instances to change it
var currentPage = welcomePage

type Pages interface {
	Init(int, int)
	Update(tea.Model, tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type SampleConfigUI struct {
	pages []Pages
}

func NewSampleConfigUI() SampleConfigUI {
	w := NewWelcomePage()
	p := NewPluginPage()

	var pages []Pages
	pages = append(pages, &w)
	pages = append(pages, &p)

	return SampleConfigUI{pages: pages}
}

func (s SampleConfigUI) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, tea.EnableMouseCellMotion)
}

func (s SampleConfigUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) { //nolint:revive
	case tea.WindowSizeMsg:
		if !ready {
			ready = true
			for _, p := range s.pages {
				p.Init(msg.Width, msg.Height)
			}
		}
	}

	model, cmd := s.pages[currentPage].Update(s, msg)
	return model, cmd
}

func (s SampleConfigUI) View() string {
	if !ready {
		return "\n  Initializing..."
	}
	return s.pages[currentPage].View()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
