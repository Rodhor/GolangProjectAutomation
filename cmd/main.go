package main

import (
	"fmt"
	"log"
	"projectAutomation/internal/pkg/parser"
	"projectAutomation/internal/pkg/project"
	"projectAutomation/internal/pkg/ui"
)

func main() {
	// retrieve a list of languages including possible packages from the embedded files
	langs, errors := parser.RetrieveEmbeddedLanguages()
	if len(errors) > 0 {
		log.Print("error retrieveing the embedded languages. Please check the folderstructure and yamlfiles.")
	}

	// Create a prointer to the project which will hold all information and run the needed functions.
	p := &project.Project{}

	ui.MainUI(langs, p)

	fmt.Println(p.Name)
	fmt.Println(p.Language.Commands)
	fmt.Println(p.Packages)
	fmt.Println(p.Structure.ID)
}
