package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	projectName State = iota
	languageSelection
	libarySelection
	libarySubDir
	SummaryAndConfirmation
)

type mainModel struct{}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m mainModel) View() string {
	return "Hello, World!"
}

func (m mainModel) New() *mainModel {
	return &mainModel{}
}

func main() {
	p := tea.NewProgram(mainModel{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
