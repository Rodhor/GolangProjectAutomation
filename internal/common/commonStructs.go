package common

type RunTime int

const (
	Init RunTime = iota
	BeforeFolderCreation
	BeforePackageInstallation
	AfterPackageInstallation
	End
)

type FileStructure struct {
	ID       string                     `ymal:"-"`
	Contents map[string]FileOrDirectory `yaml:"contents"`
	Commands map[string]Command         `yaml:"commands, omitempty"`
}

type FileOrDirectory struct {
	Content  *string                    `yaml:"content, omitempty"`
	Children map[string]FileOrDirectory `yaml:", omitempty"`
}

type Command struct {
	Cmd     string `yaml:"cmd"`
	RunTime string `yaml:"runtime"`
}
