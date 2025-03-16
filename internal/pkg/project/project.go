package project

import (
	"projectAutomation/internal/common"
	"projectAutomation/internal/config"
	"strings"
)

type Project struct {
	Name            string
	Language        *config.Language
	Structure       common.FileStructure
	Packages        []config.Package
	RenderingOption common.Confirmation
	RootDir         string
}

func (p *Project) correctRuntime() {
	for _, cmd := range p.Language.Commands {
		switch strings.ToLower(cmd.RunTime) {
		case "init":
			cmd.ActualRunTime = common.Init
		case "beforefoldercreation":
			cmd.ActualRunTime = common.BeforeFolderCreation
		case "beforepackageinstallation":
			cmd.ActualRunTime = common.BeforePackageInstallation
		case "afterpackageinstallation":
			cmd.ActualRunTime = common.AfterPackageInstallation
		case "end":
			cmd.ActualRunTime = common.End
		default:
			cmd.ActualRunTime = common.End
		}
	}
}

func (p *Project) createProject() error {
	return nil
}
