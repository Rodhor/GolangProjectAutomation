package config

type Shell struct {
	Shell   string
	ArgFlag string
}

type PlaceholderArg string

const (
	ArgProjectName     PlaceholderArg = "{{project_name}}"
	ArgProjectPath     PlaceholderArg = "{{project_path}}"
	ArgProjectLanguage PlaceholderArg = "{{project_language}}"
)
