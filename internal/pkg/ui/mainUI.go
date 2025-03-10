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
	var folderstructure string
	var packages []string
	var createPackageStructure bool

	result := make(map[string]any)

	group := huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose your desired Folderstructure").
			Value(&folderstructure).
			OptionsFunc(func() []huh.Option[string] {
				opts := fetchFolderStructures(lang)
				if len(opts) == 0 {
					return []huh.Option[string]{huh.NewOption("None", "none")}
				}
				return opts
			}, &folderstructure),
		// Only add package selection if options are available.
		func() huh.Field {
			opts := fetchPackageOptions(lang)
			if len(opts) > 0 {
				return huh.NewMultiSelect[string]().
					Title("Choose your desired Packages").
					Value(&packages).
					OptionsFunc(func() []huh.Option[string] {
						return opts
					}, nil)
			}
			return huh.NewInput().Title("No packages available").Value(new(string))
		}(),
		huh.NewSelect[bool]().
			Title("Should the folderstructure for each package be rendered?").
			Options(
				huh.NewOption("Yes", true),
				huh.NewOption("No", false),
			).
			Value(&createPackageStructure),
	)

	form := huh.NewForm(group)
	err := form.Run()
	if err != nil {
		return nil, err
	}

	result["Folderstructure"] = folderstructure
	result["Packages"] = packages
	result["Create Package Structure"] = createPackageStructure

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

func main() {
	// For demonstration, assume availableLanguages is pre-populated.
	availableLanguages := []config.Language{
		{ID: "golang", Name: "Go"},
		{ID: "python", Name: "Python"},
	}
	MainUI(availableLanguages)
}
