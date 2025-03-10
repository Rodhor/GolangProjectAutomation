package main

import (
	"fmt"
	"log"
	"projectAutomation/internal/pkg/parser"
	"projectAutomation/internal/pkg/ui"
)

// ==================================================
// Lipgloss Styling and Views (Placeholders)
// ==================================================

// TODO: Add lipgloss styling and view helper functions here.

// ==================================================
// Main Program Entry Point
// ==================================================

var testDir = "/home/rodhor/Documents/test/"

func main() {
	langs, errors := parser.RetrieveEmbeddedLanguages()
	if len(errors) > 0 {
		log.Print("Damn")
	}
	lang := langs[0]
	for _, fs := range lang.FileStructure {
		fmt.Println(fs.ID)
	}

	ui.MainUI(langs)
}
