package scaffold

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Handle Merging of structures
func MergeStructures(structures ...FileStructure) FileStructure {
	merged := FileStructure{}
	dirSet := make(map[string]struct{})
	fileSet := make(map[string]struct{})

	for _, s := range structures {

		// Check subdirs
		for _, dir := range s.SubDirs {
			if _, exists := dirSet[dir]; !exists {
				dirSet[dir] = struct{}{}
				merged.SubDirs = append(merged.SubDirs, dir)
			}
		}

		// Check files
		for _, f := range s.Files {
			if _, exists := fileSet[f.Name]; !exists {
				fileSet[f.Name] = struct{}{}
				merged.Files = append(merged.Files, f)
			}
		}
	}

	return merged
}

func BuildCompleteStructure(lang Language, selectedLevel StructureLevel, selectedPackages []Package) (FileStructure, string) {
	completeStructure := lang.Scaffold.GetStructure(selectedLevel)
	finalMakefileContent := lang.Scaffold.GetMakefileContent()

	for _, pkg := range selectedPackages {
		pkgStructure := pkg.Scaffold.GetFileStructure()
		completeStructure = MergeStructures(completeStructure, pkgStructure)
		if pkg.Scaffold.GetMakeFileContent() != "" {
			finalMakefileContent += pkg.Scaffold.GetMakeFileContent()
		}
	}

	return completeStructure, finalMakefileContent
}

// CreateFolders creates all directories specified in the FileStructure.
func CreateFolders(fs FileStructure, projectRoot string) {
	for _, dir := range fs.SubDirs {
		finalPath := filepath.Join(projectRoot, dir)
		if err := os.MkdirAll(finalPath, os.ModePerm); err != nil {
			log.Fatalf("Error creating directory %s: %v", finalPath, err)
		}
	}
}

// CreateFiles creates all files specified in the FileStructure.
func CreateFiles(fs FileStructure, projectRoot string) {
	for _, file := range fs.Files {
		finalDir := filepath.Join(projectRoot, file.PlacementDir)
		// Ensure the directory exists before creating the file.
		if err := os.MkdirAll(finalDir, os.ModePerm); err != nil {
			log.Fatalf("Error creating directory for file %s: %v", file.Name, err)
		}
		filePath := filepath.Join(finalDir, file.Name)
		if err := os.WriteFile(filePath, []byte(file.Content), 0644); err != nil {
			log.Fatalf("Error writing file %s: %v", filePath, err)
		}
	}
}

// CreateMakefile writes the provided Makefile content.
func CreateMakefile(makefileContent string, projectRoot string) {
	makefilePath := filepath.Join(projectRoot, "Makefile")
	if err := os.WriteFile(makefilePath, []byte(makefileContent), 0644); err != nil {
		log.Fatalf("Error writing Makefile: %v", err)
	}
}

// RunCommands executes commands based on their runtime stage.
func RunCommands(cmds []Command, projectRoot string, runtime Runtime) {
	for _, command := range cmds {
		if command.RunTime == runtime {
			cmd := exec.Command("sh", "-c", command.Cmd)
			cmd.Dir = projectRoot
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Command '%s' failed: %v, output: %s", command.Cmd, err, string(output))
			} else {
				log.Printf("Command '%s' succeeded, output: %s", command.Cmd, string(output))
			}
		}
	}
}

// InstallPackages runs installation commands for selected packages.
func InstallPackages(packages []Package, projectRoot string) {
	for _, pkg := range packages {
		if pkg.InstallCmd != "" {
			cmd := exec.Command("sh", "-c", pkg.InstallCmd)
			cmd.Dir = projectRoot
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Failed to install package '%s': %v, output: %s", pkg.Name, err, string(output))
			} else {
				log.Printf("Package '%s' installed successfully, output: %s", pkg.Name, string(output))
			}
		}
	}
}

// SetupProjectStructure creates the full project structure and runs necessary commands.
func SetupProjectStructure(projectRoot string, fs FileStructure, makefileContent string, commands []Command, selectedPackages []Package) {
	// Run initial commands
	RunCommands(commands, projectRoot, Init)

	// Create folders and files
	CreateFolders(fs, projectRoot)
	CreateFiles(fs, projectRoot)

	// Create Makefile
	CreateMakefile(makefileContent, projectRoot)

	RunCommands(commands, projectRoot, BeforePackageInstall)

	// Install packages
	InstallPackages(selectedPackages, projectRoot)

	// Run remaining commands
	RunCommands(commands, projectRoot, AfterPackageInstall)
	RunCommands(commands, projectRoot, End)
}
