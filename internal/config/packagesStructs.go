package config

import "projectAutomation/internal/common"

type Package struct {
	ID              string                `yaml:"id"`
	Name            string                `yaml:"name"`
	LanguageID      string                `yaml:"languageID"`
	FileStructure   *common.FileStructure `yaml:"filestructure, omitempty"`
	MakefileContent *string               `ymal:"makefileContent, omitempty"`
}
