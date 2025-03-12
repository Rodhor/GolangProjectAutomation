package ui

import (
	"errors"
	"projectAutomation/internal/config"

	"github.com/charmbracelet/huh"
)

func selectLanguage(langs []config.Language) (FirstSelection, error) {
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
				}, nil),
		),
	)
	err := form.Run()
	if err != nil {
		return FirstSelection{}, err
	}
	result := FirstSelection{
		name:     projectName,
		language: languageID,
	}

	return result, nil
}
