package cli

import tea "github.com/charmbracelet/bubbletea"

type CliTool struct {
	windows   map[string]Windows
	window    string
	statusbar *StatusBar
}

func (m *CliTool) Init() tea.Cmd {
	return nil
}

func (m *CliTool) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.windows[m.window].Update(msg)
}

func (m *CliTool) View() string {
	return m.windows[m.window].View()
}

func InitCliTool() *CliTool {
	c := &CliTool{}

	c.window = "MAINMENU"
	c.windows = map[string]Windows{
		"MAINMENU":  InitMenu(),
		"SMARTSYNC": InitSmartSync(),
	}
	c.statusbar = &StatusBar{
		mode:        "NORMAL",
		message:     "Hello Shreyas, more configurations in march.",
		comment:     "Sleep well",
		encoding:    "UTF-8",
		coordinates: "[_,_]",
	}

	return c
}
