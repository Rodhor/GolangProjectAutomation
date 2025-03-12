package config

import "projectAutomation/internal/common"

type Project struct {
	ProjectName            string
	ProjectLanguage        string
	ProjectStructure       string
	ProjectPackages        []string
	ProjectRenderingOption common.Confirmation
	ProjectDir             string
}
