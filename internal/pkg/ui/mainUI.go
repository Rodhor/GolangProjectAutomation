package ui

import (
	"fmt"
	"log"
	"projectAutomation/internal/config"
)

func MainUI(langs []config.Language) {
	var p config.Project
	firstSelection, err := selectLanguage(langs)
	if err != nil {
		log.Fatal(err)
	}
	p.ProjectName = firstSelection.name
	p.ProjectLanguage = firstSelection.language

	var selectedLang *config.Language
	for i, lang := range langs {
		if lang.ID == p.ProjectLanguage {
			selectedLang = &langs[i]
			break
		}
	}
	if selectedLang == nil {
		log.Fatalf("Language with ID %s not found", p.ProjectLanguage)
	}

	languagespecificSelection, err := selectLanguageSpecificOptions(*selectedLang)
	if err != nil {
		log.Fatal(err)
	}
	p.ProjectStructure = languagespecificSelection.folderStructure
	p.ProjectPackages = languagespecificSelection.packages
	pkgSelected := len(p.ProjectPackages) != 0
	overallConfirmation, err := getConfirmationToExecute(pkgSelected)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(overallConfirmation)
}
