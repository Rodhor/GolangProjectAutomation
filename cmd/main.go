package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// ==================================================
// Lipgloss Styling and Views (Placeholders)
// ==================================================

// TODO: Add lipgloss styling and view helper functions here.

// ==================================================
// Domain Models and Data Structures
// ==================================================

// File represents a file with its type, content, and placement directory.
type File struct {
	FileType     string
	Content      string
	PlacementDir string
}

// Package represents a package (e.g., Pandas, bubbletea) that can be installed
// for a programming language.
type Package struct {
	Name       string
	SubDir     []string
	SubFiles   []File
	InstallCmd string
}

// Language represents a programming language (e.g., Python, Go)
// Depending on the language, it may have different packages and files.
type Language struct {
	Name     string
	SubDir   []string
	subFiles []File
	Packages []Package
}

// FetchPossibleLanguages returns a list of available programming languages.
func FetchAllLanguages() []Language {
	return []Language{}
}

// FetchPossiblePackages returns a list of available packages for a given language.
func FetchPossiblePackages(languageSelection string) []Package {
	return []Package{}
}

// ==================================================
// Bubble Tea Model and Application States
// ==================================================

// State is an enumeration that represents the current view/state.
type State int

const (
	projectName State = iota
	languageSelection
	packageSelection
	packageSubDir
	SummaryAndConfirmation
)

// mainModel is the Bubble Tea model for the main application.
type mainModel struct {
	state State
}

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
	p := tea.NewProgram(mainModel{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
