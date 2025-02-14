package scaffold

type StructureLevel int

const (
	LevelNone StructureLevel = iota
	LevelBasic
	LevelWebDevelopment
	LevelAdvanced
	LevelProduction
)

type Runtime int

const (
	Init Runtime = iota
	BeforePackageInstall
	AfterPackageInstall
	End
)

type FileStructure struct {
	SubDirs []string
	Files   []File
}

type File struct {
	Name         string
	Content      string
	PlacementDir string
}

type Language struct {
	Name     string
	Scaffold LanguageScaffold
	Packages []Package
	Commands []Command
}

type Command struct {
	Cmd     string
	RunTime Runtime
}

type Package struct {
	Name       string
	InstallCmd string
	Scaffold   PackageScaffold
}
