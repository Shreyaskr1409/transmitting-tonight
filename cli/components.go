package cli

import tea "github.com/charmbracelet/bubbletea"

type Windows interface {
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type StatusBar struct {
	mode        string
	encoding    string
	comment     string
	message     string
	coordinates string
}
