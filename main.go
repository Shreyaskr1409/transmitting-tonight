package main

import (
	// "fmt"
	"log"
	// "os"
	"transmitting-tonight/cli"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mm := cli.InitCliTool()
	p := tea.NewProgram(mm, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
