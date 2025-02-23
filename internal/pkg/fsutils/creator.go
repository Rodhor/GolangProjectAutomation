package fsutils

import (
	"fmt"
	"os"
	"path/filepath"
	"projectAutomation/internal/common"
)

func CreateFiletree(rootPath string, fs map[string]common.FileOrDirectory) error {
	// Loop through all entries within a "folder" in a folderstructure
	for name, node := range fs {
		path := filepath.Join(rootPath, name)

		// If a child exists it is a directory, create the directory recursively
		// until all subdirectories and folders are created
		if node.Children != nil {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", path, err)
			}

			// Recursively create subdirectories and files
			if node.Children != nil {
				if err := CreateFiletree(path, node.Children); err != nil {
					return err
				}
			}

			// If it is not a folder, create the file
		} else {
			file, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %v", path, err)
			}
			if node.Content != nil {
				if _, err := file.WriteString(*node.Content); err != nil {
					file.Close()
					return fmt.Errorf("failed to write content to file %s, %v", path, err)
				}
			}
			file.Close()
		}
	}
	return nil
}
