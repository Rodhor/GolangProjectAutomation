package common

type RunTime int

const (
	Init RunTime = iota
	BeforeFolderCreation
	BeforePackageInstallation
	AfterPackageInstallation
	End
)

type Confirmation int

const (
	OnlyLanguage Confirmation = iota
	RenderFully
	WithoutCommands
	Cancel
)

type FileStructure struct {
	ID       string                     `ymal:"-"`
	Contents map[string]FileOrDirectory `yaml:"contents"`
	Commands map[string]Command         `yaml:"commands,omitempty"`
}

type FileOrDirectory struct {
	Type     string                     `yaml:"type,omitempty"`
	Content  *string                    `yaml:"content,omitempty"`
	Children map[string]FileOrDirectory `yaml:"children,omitempty"`
}

type Command struct {
	Cmd           string `yaml:"cmd"`
	RunTime       string `yaml:"runtime"`
	ActualRunTime RunTime
}
