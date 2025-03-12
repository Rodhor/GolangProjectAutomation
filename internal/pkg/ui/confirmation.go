package ui

import (
	"projectAutomation/internal/common"

	"github.com/charmbracelet/huh"
)

func getConfirmationToExecute(pkgSelected bool) (common.Confirmation, error) {
	var renderSelection common.Confirmation

	if pkgSelected {
		huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[common.Confirmation]().
					Title("How do you wish to proceed with your selection?").
					Value(&renderSelection).
					Options(
						huh.NewOption("Create Languagespecific Folderstructure, and install packages (Packagespecific Folderstructure will not be implemented)", common.OnlyLanguage),
						huh.NewOption("Create Language- and packagespecific Folderstructure, and install packages", common.RenderFully),
						huh.NewOption("Create Languagespecific Folderstructure, without running any commands", common.WithoutCommands),
						huh.NewOption("Cancel and exit program.", common.Cancel),
					),
			)).Run()
	} else {
		huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[common.Confirmation]().
					Title("How do you wish to proceed with your selection?").
					Value(&renderSelection).
					Options(
						huh.NewOption("Create Languagespecific Folderstructure", common.OnlyLanguage),
						huh.NewOption("Create Languagespecific Folderstructure, without running any commands", common.WithoutCommands),
						huh.NewOption("Cancel and exit program.", common.Cancel),
					),
			)).Run()
	}

	return renderSelection, nil
}
