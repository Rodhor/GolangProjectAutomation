package runner

import (
	"projectAutomation/internal/config"
	"projectAutomation/internal/pkg/project"
	"regexp"
	"runtime"
)

func GetShell() config.Shell {
	if runtime.GOOS == "windows" {
		return config.Shell{Shell: "cmd", ArgFlag: "/c"}
	}
	return config.Shell{Shell: "sh", ArgFlag: "-c"}
}

func ResolvePlaceholder(p config.PlaceholderArg, project *project.Project) string {
	switch p {

	case config.ArgProjectName:
		return GetProjectName(project)

	case config.ArgProjectPath:
		return GetProjectPath(project)

	case config.ArgProjectLanguage:
		return GetProjectLanguage(project)

	default:
		return "{{" + string(p) + "}}"
	}
}

func AdjustDynamicCommands(cmd string, project *project.Project) string {
	re := regexp.MustCompile(`\{\{\s*(\w+)\s*\}\}`)
	return re.ReplaceAllStringFunc(cmd, func(match string) string {
		// Use FindStringSubmatch to extract the variable name without extra spaces.
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		key := submatches[1]
		p := config.PlaceholderArg(key)
		return ResolvePlaceholder(p, project)
	})
}

func GetProjectName(p *project.Project) string {
	return p.Name
}

func GetProjectPath(p *project.Project) string {
	return p.RootDir
}

func GetProjectLanguage(p *project.Project) string {
	return p.Language.Name
}
