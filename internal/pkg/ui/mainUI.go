package ui

import (
	"errors"
	"fmt"
	"log"
	"projectAutomation/internal/config"

	"github.com/charmbracelet/huh"
)

// Step 1: Language Selection UI
func selectLanguage(langs []config.Language) (map[string]string, error) {
	var languageID string
	var projectName string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What should the name of your Project be?").
				Value(&projectName).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("please provide a project name")
					}
					return nil
				}),

			huh.NewSelect[string]().
				Title("Choose a language").
				Value(&languageID).
				OptionsFunc(func() []huh.Option[string] {
					var opts []huh.Option[string]
					for _, lang := range langs {
						opts = append(opts, huh.NewOption(lang.Name, lang.ID))
					}
					return opts
				}, &languageID),
		),
	)
	err := form.Run()
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	result["Name"] = projectName
	result["Language"] = languageID

	return result, nil
}

// Step 2: Language-Specific Options (Folder structure, packages, etc.)
func selectLanguageSpecificOptions(lang config.Language) (map[string]any, error) {
	var (
		folderstructure        string
		packages               []string
		createPackageStructure bool
	)

	folderstructureOptions := fetchFolderStructures(lang)

	fields := []huh.Field{
		huh.NewSelect[string]().
			Title("Choose your desired Folderstructure").
			Value(&folderstructure).
			OptionsFunc(func() []huh.Option[string] {
				return folderstructureOptions
			}, &folderstructure),
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

		fields = append(fields, huh.NewConfirm().
			Title("Would you like to create Packagespecific folderstructures and files?").
			Value(&createPackageStructure),
		)
	}

	group := huh.NewGroup(fields...)

	form := huh.NewForm(group)
	err := form.Run()
	if err != nil {
		return nil, err
	}

	result := map[string]any{
		"Folderstructure":          folderstructure,
		"Packages":                 packages,
		"Create Package Structure": createPackageStructure,
	}

	return result, nil
}

// Step 3: Overall confirmation and running the programm
func getConfirmationToExecute() {
}

func MainUI(langs []config.Language) {
	firstSelection, err := selectLanguage(langs)
	if err != nil {
		log.Fatal(err)
	}

	var selectedLang *config.Language
	for i, lang := range langs {
		if lang.ID == firstSelection["Language"] {
			selectedLang = &langs[i]
			break
		}
	}
	if selectedLang == nil {
		log.Fatalf("Language with ID %s not found", firstSelection["Language"])
	}

	fmt.Printf("Selected Language: %s\n", selectedLang.Name)

	secondSelection, err := selectLanguageSpecificOptions(*selectedLang)
	if err != nil {
		log.Fatal(err)
	}

	// Continue with further processing as needed.
	log.Println("Language-specific options selected successfully.")
	log.Println(secondSelection)
}

func fetchFolderStructures(lang config.Language) []huh.Option[string] {
	var options []huh.Option[string]
	for _, fs := range lang.FileStructure {
		options = append(options, huh.NewOption(fs.ID, fs.ID))
	}
	return options
}

func fetchPackageOptions(lang config.Language) []huh.Option[string] {
	var options []huh.Option[string]
	if lang.PackageIDs == nil {
		return options
	}
	for _, pkg := range *lang.PackageIDs {
		options = append(options, huh.NewOption(pkg, pkg))
	}
	return options
}
