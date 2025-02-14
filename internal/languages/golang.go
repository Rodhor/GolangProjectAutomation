package languages

import (
	s "projectAutomation/internal/pkg/scaffold"
)

type GoScaffold struct{}

var BasicStructure = s.FileStructure{
	SubDirs: []string{
		"cmd/app",
		"internal/pkg/fsutils",
		"pkg",
		"docs",
	},
	Files: []s.File{
		{
			Name:         "main.go",
			Content:      "package main",
			PlacementDir: "cmd/app/",
		},
		{
			Name:         "fsutils.go",
			Content:      "",
			PlacementDir: "internal/pkg/fsutils/",
		},
		{
			Name:         ".gitignore",
			Content:      "",
			PlacementDir: "./",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "docs/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "internal/pkg/fsutils/",
		},
	},
}

var WebDevelopmentStructure = s.FileStructure{
	SubDirs: []string{
		"cmd/app",
		"internal/pkg/fsutils",
		"pkg",
		"api",
		"web",
		"examples",
		"docs",
	},
	Files: []s.File{
		{
			Name:         "main.go",
			Content:      "package main",
			PlacementDir: "cmd/app/",
		},
		{
			Name:         "fsutils.go",
			Content:      "package fsutils",
			PlacementDir: "internal/pkg/fsutils/",
		},
		{
			Name:         ".gitignore",
			Content:      "",
			PlacementDir: "./",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "docs/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "internal/pkg/fsutils/",
		},
	},
}

var AdvancedStructure = s.FileStructure{
	SubDirs: []string{
		"cmd/app",
		"internal/pkg/fsutils",
		"internal/pkg/types",
		"test",
		"config",
		"docs",
	},
	Files: []s.File{
		{
			Name:         "main.go",
			Content:      "package main",
			PlacementDir: "cmd/app/",
		},
		{
			Name:         "fsutils.go",
			Content:      "package fsutils",
			PlacementDir: "internal/pkg/fsutils/",
		},
		{
			Name:         "types.go",
			Content:      "package types",
			PlacementDir: "internal/pkg/types/",
		},
		{
			Name:         ".gitignore",
			Content:      "",
			PlacementDir: "./",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "internal/pkg/types/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "docs/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "internal/pkg/fsutils/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "config/",
		},
	},
}

var ProductionStructure = s.FileStructure{
	SubDirs: []string{
		"cmd/app",
		"internal/pkg/fsutils",
		"internal/service",
		"pkg",
		"configs",
		"deploy",
		"build",
		"scripts",
		"test",
		"docs",
		"logs",
	},
	Files: []s.File{
		{
			Name:         "main.go",
			Content:      "package main",
			PlacementDir: "cmd/app/",
		},
		{
			Name:         "fsutils.go",
			Content:      "package fsutils",
			PlacementDir: "internal/pkg/fsutils/",
		},
		{
			Name:         "baseModels.go",
			Content:      "package types",
			PlacementDir: "internal/pkg/types/",
		},
		{
			Name:         "service.go",
			Content:      "package service",
			PlacementDir: "internal/service/",
		},
		{
			Name:         "config.yaml",
			Content:      "",
			PlacementDir: "configs/",
		},
		{
			Name:         "prod.yaml",
			Content:      "",
			PlacementDir: "configs/",
		},
		{
			Name:         "release.sh",
			Content:      "",
			PlacementDir: "scripts/",
		},
		{
			Name:         "migrate.sh",
			Content:      "",
			PlacementDir: "scripts/",
		},
		{
			Name:         ".gitignore",
			Content:      "",
			PlacementDir: "./",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "internal/pkg/types/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "docs/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "internal/pkg/fsutils/",
		},
		{
			Name:         "README.md",
			Content:      "",
			PlacementDir: "config/",
		},
	},
}

func (g GoScaffold) SupportedStructureLevels() []s.StructureLevel {
	return []s.StructureLevel{
		s.LevelBasic,
		s.LevelWebDevelopment,
		s.LevelAdvanced,
		s.LevelProduction,
	}
}

func (g GoScaffold) GetStructure(level s.StructureLevel) s.FileStructure {
	switch level {
	case s.LevelBasic:
		return BasicStructure
	case s.LevelWebDevelopment:
		return WebDevelopmentStructure
	case s.LevelAdvanced:
		return AdvancedStructure
	case s.LevelProduction:
		return ProductionStructure
	default:
		return s.FileStructure{}
	}
}

func (g GoScaffold) GetMakeFileContent() string {
	return ""
}

func (g GoScaffold) GetCommands() []s.Command {
	return []s.Command{
		{
			Cmd:     "go mod init",
			RunTime: s.BeforePackageInstall,
		},
		{
			Cmd:     "git init",
			RunTime: s.End,
		},
		{
			Cmd:     "git add .",
			RunTime: s.End,
		},
		{
			Cmd:     "git commit -m 'Initial commit'",
			RunTime: s.End,
		},
	}
}
