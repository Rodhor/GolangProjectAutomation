package ui

import (
	"projectAutomation/internal/config"

	"github.com/charmbracelet/huh"
)

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
