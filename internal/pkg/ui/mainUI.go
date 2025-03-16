package ui

import (
	"log"
	"projectAutomation/internal/config"
	"projectAutomation/internal/pkg/project"
)

func MainUI(langs []config.Language, p *project.Project) {
	firstSelection, err := selectLanguage(langs)
	if err != nil {
		log.Fatal(err)
	}
	p.Name = firstSelection.name
	languageSelection := firstSelection.language
	for i, lang := range langs {
		if lang.ID == languageSelection {
			p.Language = &langs[i]
			break
		}
	}
	if p.Language == nil {
		log.Fatalf("Language with ID %s not found", languageSelection)
	}

	languagespecificSelection, err := selectLanguageSpecificOptions(*p.Language)
	if err != nil {
		log.Fatal(err)
	}
	for _, folderstruct := range p.Language.FileStructure {
		if folderstruct.ID == languagespecificSelection.folderStructure {
			p.Structure = folderstruct
		}
	}

	var packages []config.Package
	for _, pkg := range *p.Language.LanguagePackages {
		for _, pkgid := range languagespecificSelection.packages {
			if pkg.ID == pkgid {
				packages = append(packages, pkg)
			}
		}
	}

	p.Packages = packages
}
