package config

import "projectAutomation/internal/common"

type Language struct {
	ID              string                 `yaml:"id"`
	Name            string                 `yaml:"name"`
	FileStructure   []common.FileStructure `yaml:"filestructure, omitempty"`
	MakefileNeeded  *bool                  `yaml:"makefileNeeded, omitempty"`
	MakeFileContent *string                `ymal:"makefileContent, omitempty"`
	Command         *[]common.Command      `yaml:"commands, omitempty"`
	PackageIDs      *[]string              `yaml:"packageIDs, omitempty"`
}
