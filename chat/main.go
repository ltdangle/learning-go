
package main

import (
	"strings"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	messages   []string
	input      textarea.Model
	msgViewport viewport.Model
}

func initialModel() *model {
	textarea := textarea.New()
	textarea.Placeholder = "Type your message and press Enter..."
	textarea.Focus()

	msgViewport := viewport.Model{}
	msgViewport.YPosition = 0

	return &model{
		messages:   []string{},
		input:      textarea,
		msgViewport: msgViewport,
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.input.Value() != "" {
				m.messages = append(m.messages, m.input.Value())
				m.input.SetValue("")
			}
		default:
			input, cmd := m.input.Update(msg)
			m.input = input
			cmds = append(cmds, cmd)
		}
	case tea.WindowSizeMsg:
		m.msgViewport.Width = msg.Width
		m.msgViewport.Height = msg.Height - 3 // reserve some space for the input field
	}

	m.msgViewport.SetContent(strings.Join(m.messages, "\n"))

	if len(cmds) > 0 {
		cmd = tea.Batch(cmds...)
	}

	return m, cmd
}

func (m *model) View() string {
	var b strings.Builder

	b.WriteString(m.msgViewport.View())
	b.WriteString("\n\n")
	b.WriteString(m.input.View())

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
