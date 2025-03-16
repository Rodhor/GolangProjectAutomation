package parser

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"projectAutomation/internal/config"
)

//go:embed languages/*
var fsys embed.FS

func RetrieveEmbeddedLanguages() ([]config.Language, []error) {
	var langs []config.Language
	var errs []error

	entries, err := fs.ReadDir(fsys, "languages")
	if err != nil {
		return nil, []error{fmt.Errorf("error reading languages directory: %w", err)}
	}

	// Create an instance of the YAML parser to parse the files
	parser := EmbeddedYamlParser{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		langDir := filepath.Join("languages", entry.Name())
		langYamlPath := ""

		langFiles, err := fs.ReadDir(fsys, langDir)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading directory %s: %w", langDir, err))
			continue
		}

		for _, fileEntry := range langFiles {
			if !fileEntry.IsDir() && parser.CheckFileType(fileEntry.Name()) {
				langYamlPath = filepath.Join(langDir, fileEntry.Name())
				break
			}
		}

		if langYamlPath == "" {
			errs = append(errs, fmt.Errorf("no YAML file found in %s", langDir))
			continue
		}

		lang, err := parser.ParseLanguage(fsys, langYamlPath)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to parse language YAML file %s: %w", langYamlPath, err))
			continue
		}

		pkgDir := filepath.Join(langDir, "packages")
		if info, err := fs.Stat(fsys, pkgDir); err == nil && info.IsDir() {
			pkgEntries, err := fs.ReadDir(fsys, pkgDir)
			if err != nil {
				errs = append(errs, fmt.Errorf("error reading the package directory %s: %w", pkgDir, err))
			} else {
				var pkgs []config.Package
				for _, pkgEntry := range pkgEntries {
					if !pkgEntry.IsDir() && parser.CheckFileType(pkgEntry.Name()) {
						pkgPath := filepath.Join(pkgDir, pkgEntry.Name())
						pkg, err := parser.ParsePackagesForLanguage(fsys, pkgPath)
						if err != nil {
							errs = append(errs, fmt.Errorf("error parsing package YAML file %s: %w", pkgPath, err))
							continue
						}

						pkgs = append(pkgs, *pkg)
					}
				}
				lang.LanguagePackages = &pkgs
			}
		}
		langs = append(langs, *lang)
	}
	return langs, errs
}
