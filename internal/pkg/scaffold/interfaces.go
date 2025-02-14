package scaffold

type LanguageScaffold interface {
	SupportedStructureLevels() []StructureLevel
	GetStructure(level StructureLevel) FileStructure
	GetMakefileContent() string
	GetCommands() []Command
}

type PackageScaffold interface {
	GetFileStructure() FileStructure
	GetMakeFileContent() string
}
