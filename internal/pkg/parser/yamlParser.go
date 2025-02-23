package parser

import (
	"fmt"
	"io/fs"
	"projectAutomation/internal/config"
	"strings"

	"gopkg.in/yaml.v3"
)

type EmbeddedYamlParser struct{}

func (p EmbeddedYamlParser) ParseLanguage(fsys fs.FS, path string) (*config.Language, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file %s: %w", path, err)
	}

	var lang config.Language
	if err := yaml.Unmarshal(data, &lang); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML from %s: %w", path, err)
	}
	return &lang, nil
}

func (p EmbeddedYamlParser) ParsePackagesForLanguage(fsys fs.FS, path string) (*config.Package, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file %s, %w", path, err)
	}

	var pkg config.Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML from %s: %w", path, err)
	}
	return &pkg, nil
}

func (p EmbeddedYamlParser) CheckFileType(name string) bool {
	lower := strings.ToLower(name)
	return strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml")
}
