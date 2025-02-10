package fsutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetHomeDirectory retrieves the current user's home directory.
func GetHomeDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error fetching home directory: %w", err)
	}
	return homeDir, nil
}

// checkIfDirectoryExists checks if a directory exists (case-insensitively)
// in the provided parent path. It returns the actual name of the directory if found.
func checkIfDirectoryExists(parentPath, targetDir string) (string, bool) {
	entries, err := os.ReadDir(parentPath)
	if err != nil {
		return "", false
	}
	targetLower := strings.ToLower(targetDir)
	for _, entry := range entries {
		if entry.IsDir() {
			if strings.ToLower(entry.Name()) == targetLower {
				return entry.Name(), true
			}
		}
	}
	return "", false
}

// EnsureDirectory ensures that a directory exists at the given path (constructed from
// a parent directory and a target directory name). If it does not exist (considering case),
// it creates the directory. It returns the full path to the directory.
func EnsureDirectory(parentPath, targetDir string) (string, error) {
	// First, check if the directory already exists (case-insensitively)
	actualDirName, exists := checkIfDirectoryExists(parentPath, targetDir)
	var fullPath string

	if exists {
		fullPath = filepath.Join(parentPath, actualDirName)
	} else {
		// If it doesn't exist, create it
		fullPath = filepath.Join(parentPath, targetDir)
		if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", fullPath, err)
		}
	}
	return fullPath, nil
}

// GetProjectsPath retrieves the final Projects path by ensuring that
// the "Documents" and "Projects" directories exist under the home directory.
func GetProjectsPath() (string, error) {
	// Retrieve the home directory
	homeDir, err := GetHomeDirectory()
	if err != nil {
		return "", err
	}

	// Ensure the "Documents" directory exists within the home directory.
	documentsPath, err := EnsureDirectory(homeDir, "Documents")
	if err != nil {
		return "", fmt.Errorf("failed to ensure Documents directory: %w", err)
	}

	// Ensure the "Projects" directory exists within the Documents directory.
	projectsPath, err := EnsureDirectory(documentsPath, "Projects")
	if err != nil {
		return "", fmt.Errorf("failed to ensure Projects directory: %w", err)
	}

	return projectsPath, nil
}
