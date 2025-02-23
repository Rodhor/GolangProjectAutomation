package main

import (
	"log"
	"projectAutomation/internal/pkg/fsutils"
	"projectAutomation/internal/pkg/parser"

	tea "github.com/charmbracelet/bubbletea"
)

// ==================================================
// Lipgloss Styling and Views (Placeholders)
// ==================================================

// TODO: Add lipgloss styling and view helper functions here.

// ==================================================
// Bubble Tea Model and Application States
// ==================================================

// mainModel is the Bubble Tea model for the main application.
type mainModel struct{}

// Init is the initial command for the mainModel.
func (m mainModel) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages and updates the model accordingly.
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

// View returns the UI as a string.
func (m mainModel) View() string {
	return "Hello, World!"
}

// ==================================================
// Main Program Entry Point
// ==================================================

func main() {
	// p := tea.NewProgram(mainModel{})
	// if _, err := p.Run(); err != nil {
	// 	panic(err)
	// }

	langs, errors := parser.RetrieveEmbeddedLanguages()
	if len(errors) > 0 {
		log.Print("Damn")
	}
	fsutils.CreateFiletree("/home/rodhor/Documents/test/", langs[0].FileStructure[0].Contents)
}
