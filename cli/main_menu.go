package cli

import (
	"fmt"

	"transmitting-tonight/cli/util"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainMenu struct {
	choices  []string
	cursor   int
	messages []string
	selected map[int]struct{}
	width    int
	height   int
	status   *StatusBar
}

func (m *MainMenu) Init() tea.Cmd {
	return nil
}

func (m *MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	options := []string{
		"Smart sync                          S",
		"Configure settings                  C",
		"Sync history                        H",
		"Quit                                Q",
	}

	msg_arr := []string{
		"Welcome to Transmitting Tonight!",
		"Transmitting Tonight is a file synchronization tool to sync files between directories on local and remote systems using SSH.",
	}
	msg_arr = append(msg_arr, options...)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case tea.KeyDown.String():
			if m.cursor+1 < len(options) {
				m.cursor++
			}
		case tea.KeyUp.String():
			if m.cursor-1 >= 0 {
				m.cursor--
			}

		case "s": // Smart sync
			// TODO
		case "c": // Configure settings
			// TODO
		case "h": // Sync history
			// TODO
		case "q":
			return m, tea.Quit

		default:
			// m.status.message = "Hello Shreyas, more configurations in march."
		}

	case tea.MouseMsg:
		x := msg.X
		y := msg.Y
		m.status.coordinates = fmt.Sprint("[", x, ",", y, "]")

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	m.messages = []string{}

	style := lipgloss.NewStyle().Bold(true).Underline(true)
	util.MsgAppendln(&m.messages, style.Render(msg_arr[0]))

	style = lipgloss.NewStyle().Width(50).Italic(true).Faint(true).AlignHorizontal(lipgloss.Center).BorderStyle(lipgloss.RoundedBorder())
	util.MsgAppendln(&m.messages, style.Render(msg_arr[1]))
	util.MsgAppend(&m.messages, "\n")
	style = lipgloss.NewStyle().Bold(true)
	util.MsgAppendln(&m.messages, style.Render("OPTIONS"))

	for i := 2; i < len(msg_arr); i++ {
		if i-2 == m.cursor {
			msg_arr[i] = fmt.Sprint("> ", msg_arr[i][:len(msg_arr[i])-3], msg_arr[i][len(msg_arr[i])-1:])
		}
		util.MsgAppendln(&m.messages, msg_arr[i])
	}

	return m, nil
}

func (m *MainMenu) View() string {
	style := lipgloss.NewStyle().
		Width(m.width-1).
		Height(m.height-1).
		Align(lipgloss.Center, lipgloss.Center)

	page := ""
	s := ""

	for i := range m.messages {
		s = fmt.Sprint(s, m.messages[i])
	}
	page = fmt.Sprint(page, style.Render(s), "\n")

	style = lipgloss.NewStyle().Background(lipgloss.Color("5")).Padding(0, 1)
	page = fmt.Sprint(page, style.Render(m.status.mode))
	style = lipgloss.NewStyle().Background(lipgloss.Color("#020202")).Padding(0, 1)
	page = fmt.Sprint(page, style.Render(m.status.message))
	style = lipgloss.NewStyle().Background(lipgloss.Color("#020202"))
	s = ""
	for i := 0; i < m.width-(len(m.status.mode)+len(m.status.comment)+len(m.status.encoding)+len(m.status.message)+len(m.status.coordinates)+(2*5)); i++ {
		s = fmt.Sprint(s, " ")
	}
	page = fmt.Sprint(page, style.Render(s))
	style = lipgloss.NewStyle().Background(lipgloss.Color("#020202")).Padding(0, 1)
	page = fmt.Sprint(page, style.Render(m.status.coordinates))
	style = lipgloss.NewStyle().Background(lipgloss.Color("5")).Padding(0, 1)
	page = fmt.Sprint(page, style.Render(m.status.encoding))
	style = lipgloss.NewStyle().Background(lipgloss.Color("9")).Padding(0, 1)
	page = fmt.Sprint(page, style.Render(m.status.comment))

	return page
}

func InitMenu() *MainMenu {
	return &MainMenu{
		selected: make(map[int]struct{}),
		messages: []string{},
		cursor:   0,
		status: &StatusBar{
			mode:        "NORMAL",
			message:     "Hello Shreyas, more configurations in march.",
			comment:     "Sleep well",
			encoding:    "UTF-8",
			coordinates: "[_,_]",
		},
	}
}

func (m *MainMenu) renderMessages() string {
	// TODO
	return ""
}

func (m *MainMenu) renderStatusBar() string {
	// TODO
	return ""
}
