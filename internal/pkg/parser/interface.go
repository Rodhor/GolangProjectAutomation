package parser

import (
	"io/fs"
	"projectAutomation/internal/config"
)

type Parser interface {
	CheckFileType(path string) bool
	ParseLanguage(fs.FS, string) (*config.Language, error)
	ParsePackagesForLangauge(fs.FS, string) (*config.Package, error)
}
