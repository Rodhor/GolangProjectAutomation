package ui

import (
	"projectAutomation/internal/config"

	"github.com/charmbracelet/huh"
)

func selectLanguageSpecificOptions(lang config.Language) (LanguageSpecificSelection, error) {
	var (
		folderstructure string
		packages        []string
	)

	folderstructureOptions := fetchFolderStructures(lang)

	fields := []huh.Field{
		huh.NewSelect[string]().
			Title("Choose your desired Folderstructure").
			Value(&folderstructure).
			OptionsFunc(func() []huh.Option[string] {
				return folderstructureOptions
			}, nil),
	}

	// If packacges are available for the language, allow user to select them
	packageOptions := fetchPackageOptions(lang)
	if len(packageOptions) > 0 {
		fields = append(fields, huh.NewMultiSelect[string]().
			Title("Choose your desired Packages").
			Value(&packages).
			OptionsFunc(func() []huh.Option[string] {
				return packageOptions
			}, nil))
	}

	group := huh.NewGroup(fields...)

	form := huh.NewForm(group)
	err := form.Run()
	if err != nil {
		return LanguageSpecificSelection{}, err
	}

	result := LanguageSpecificSelection{
		folderStructure: folderstructure,
		packages:        packages,
	}

	return result, nil
}
