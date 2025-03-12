package main

import (
	"log"
	"projectAutomation/internal/pkg/parser"
	"projectAutomation/internal/pkg/ui"
)

func main() {
	langs, errors := parser.RetrieveEmbeddedLanguages()
	if len(errors) > 0 {
		log.Print("Damn")
	}
	ui.MainUI(langs)
}
